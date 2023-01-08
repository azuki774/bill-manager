package water

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"os"
	"time"

	"go.uber.org/zap"
)

const fetcherDir = "/root/data.csv"

var jst *time.Location

func init() {
	j, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	jst = j
}

type DBRepository interface {
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

	remoteDir := w.remoteRootDir + w.Date[0:6] + "/" + w.Date + ".csv"
	err = w.Downloader.Download(ctx, fetcherDir, remoteDir)
	if err != nil {
		w.Logger.Error("failed to download CSV", zap.String("remoteDir", remoteDir), zap.Error(err))
		return err
	}

	w.Logger.Info("complete download CSV or skipped", zap.String("remoteDir", remoteDir))

	return nil
}
