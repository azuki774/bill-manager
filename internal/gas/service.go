package remix

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"time"

	"go.uber.org/zap"
)

type DBRepository interface {
	AddGasBill(record model.BillGas) (err error)
}

type FileLoader interface {
	LoadGasBillCSV(ctx context.Context, dir string) (recs []model.GasBillingCSV, err error)
}

const fetcherDir = "/root/fetcher/"

var jst *time.Location

func init() {
	j, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	jst = j
}

type GasService struct {
	Logger       *zap.Logger
	DBRepository DBRepository
	FileLoader   FileLoader
	Date         string // YYYYMMDD

}

func (g *GasService) Import(ctx context.Context) (err error) {
	if g.Date == "" {
		g.Date = time.Now().In(jst).Format("20060102")
	}
	g.Logger.Info("import gas bill start", zap.String("date", g.Date))

	remoteDir := fetcherDir + g.Date[0:6] + "/" + g.Date + ".csv" // ex. /root/fetcher/2023/202301/20230101.csv
	// Assume that CSV file exists in fetcherDir
	recs, err := g.FileLoader.LoadGasBillCSV(ctx, remoteDir)
	if err != nil {
		g.Logger.Error("failed to load CSV", zap.Error(err))
		return err
	}

	g.Logger.Info("load gas bill CSV complete", zap.String("target_date", g.Date))

	for _, rec := range recs {
		r, err := rec.NewGasDBModel()
		if err != nil {
			g.Logger.Error("failed to convert DB model", zap.Error(err))
			return err
		}
		err = g.DBRepository.AddGasBill(r)
		if err != nil {
			g.Logger.Error("failed to insert DB", zap.Error(err))
			return err
		}
	}

	g.Logger.Info("complete gas bill recorded", zap.String("target_date", g.Date))
	return nil
}
