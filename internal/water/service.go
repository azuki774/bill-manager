package water

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"os"
	"time"

	"go.uber.org/zap"
)

const fetcherCSV = "./data.csv"

var jst *time.Location

func init() {
	j, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	jst = j
}

type DBRepository interface {
	AddWaterBill(r model.BillWater) (err error)
}
type FileLoader interface {
	LoadWaterBillCSV(ctx context.Context, dir string) (recs []model.WaterBillingCSV, err error)
}
type Downloader interface {
	Download(ctx context.Context, dir string, remoteDir string) (err error)
}

type WaterService struct {
	Logger       *zap.Logger
	DBRepository DBRepository
	FileLoader   FileLoader
	Date         string // YYYYMMDD
	Downloader   Downloader

	remoteRootDir string // get csv path is ${remoteRootDir}/YYYYMM/YYYYMMDD.csv
}

func (w *WaterService) getEnvValue() {
	if w.Date == "" {
		w.Date = time.Now().In(jst).Format("20060102")
	}
	w.remoteRootDir = os.Getenv("SRC_REMOTE_DIR")
}

func (w *WaterService) Import(ctx context.Context) (err error) {
	w.getEnvValue()
	w.Logger.Info("import start", zap.String("date", w.Date))

	remoteDir := w.remoteRootDir + w.Date[0:6] + "/" + w.Date + ".csv" // ex. /root/fetcher/2023/202301/20230101.csv
	err = w.Downloader.Download(ctx, fetcherCSV, remoteDir)
	if err != nil {
		w.Logger.Error("failed to download CSV", zap.String("remoteDir", remoteDir), zap.Error(err))
		return err
	}

	w.Logger.Info("complete download CSV or skipped", zap.String("remoteDir", remoteDir))

	// Assume that CSV file exists in fetcherDir
	recs, err := w.FileLoader.LoadWaterBillCSV(ctx, fetcherCSV)
	if err != nil {
		w.Logger.Error("failed to load CSV", zap.Error(err))
		return err
	}

	for _, rec := range recs {
		r, err := rec.NewWaterDBModel()
		if err != nil {
			w.Logger.Error("failed to convert DB model", zap.Error(err))
			return err
		}
		err = w.DBRepository.AddWaterBill(r)
		if err != nil {
			w.Logger.Error("failed to insert DB", zap.Error(err))
			return err
		}
	}

	return nil
}
