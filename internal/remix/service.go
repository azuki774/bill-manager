package remix

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"time"

	"go.uber.org/zap"
)

type DBRepository interface {
	AddElectConsumption(record model.ElectConsumption) (err error)
}

type FileLoader interface {
	LoadRemixElectConsumptionCSV(ctx context.Context, filePath string) (recs []model.RemixCSV, err error)
}

const fetcherDir = "/root/fetcher/"

type Importer struct {
	Logger       *zap.Logger
	DBRepository DBRepository
	FileLoader   FileLoader
	Date         string // YYYYMMDD
}

func (i *Importer) Start(ctx context.Context) (err error) {
	if i.Date == "" {
		i.Date = time.Now().Format("20060102")
	}

	i.Logger.Info("import start", zap.String("target_date", i.Date))

	dir := fetcherDir + i.Date[0:6] + i.Date + ".csv"
	rrecs, err := i.FileLoader.LoadRemixElectConsumptionCSV(ctx, dir)
	if err != nil {
		i.Logger.Error("failed to load remix CSV", zap.Error(err))
		return err
	}

	i.Logger.Info("load remix CSV complete", zap.String("target_date", i.Date))

	for _, rrec := range rrecs {
		rec, err := rrec.ConvDBModel()
		if err != nil {
			i.Logger.Error("failed to convert to DB model", zap.Error(err))
			return err
		}
		err = i.DBRepository.AddElectConsumption(rec)
		if err != nil {
			i.Logger.Error("failed to insert record to DB", zap.Error(err))
			return err
		}
	}
	i.Logger.Info("DB insert complete")
	i.Logger.Info("import complete")
	return nil
}
