package water

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"time"

	"go.uber.org/zap"
)

const fetcherDir = "/root/fetcher/water/"

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

type WaterService struct {
	Logger       *zap.Logger
	DBRepository DBRepository
	FileLoader   FileLoader
	Date         string // YYYYMMDD

}

func (w *WaterService) Import(ctx context.Context) (err error) {
	if w.Date == "" {
		w.Date = time.Now().In(jst).Format("20060102")
	}
	w.Logger.Info("import water bill start", zap.String("date", w.Date))

	remoteDir := fetcherDir + w.Date[0:6] + "/" + w.Date + ".csv" // ex. /root/fetcher/2023/202301/20230101.csv
	// Assume that CSV file exists in fetcherDir
	recs, err := w.FileLoader.LoadWaterBillCSV(ctx, remoteDir)
	if err != nil {
		w.Logger.Error("failed to load CSV", zap.Error(err))
		return err
	}

	w.Logger.Info("load water bill CSV complete", zap.String("target_date", w.Date))

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

	w.Logger.Info("complete water bill recorded", zap.String("target_date", w.Date))
	return nil
}
