package remix

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type DBRepository interface {
	AddElectConsumption(record model.ElectConsumption) (err error)
	AddElectBill(record model.BillElect) (err error)
}

type FileLoader interface {
	LoadRemixElectConsumptionCSV(ctx context.Context, filePath string) (recs []model.RemixCSV, err error)
	LoadRemixElectBillCSV(ctx context.Context, filePath string) (recs []model.RemixBillingCSV, err error)
}

const fetcherDir = "/root/fetcher/"

type Importer struct {
	Logger       *zap.Logger
	DBRepository DBRepository
	FileLoader   FileLoader
	Date         string // YYYYMMDD
}

func (i *Importer) Start(ctx context.Context, target string) (err error) {
	switch target {
	case "consume":
		err = i.startConsume(ctx)
	default:
		err = fmt.Errorf("invalid target args")
	}

	return err
}

func (i *Importer) startConsume(ctx context.Context) (err error) {
	if i.Date == "" {
		i.Date = time.Now().Format("20060102")
	}

	i.Logger.Info("import consume start", zap.String("target_date", i.Date))

	dir := fetcherDir + i.Date[0:6] + "/" + i.Date + ".csv"
	rrecs, err := i.FileLoader.LoadRemixElectConsumptionCSV(ctx, dir)
	if err != nil {
		i.Logger.Error("failed to load remix CSV", zap.Error(err))
		return err
	}

	i.Logger.Info("load remix consume CSV complete", zap.String("target_date", i.Date))

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
	i.Logger.Info("DB insert consume complete")
	i.Logger.Info("import consume complete")
	return nil
}

func (i *Importer) startBill(ctx context.Context) (err error) {
	if i.Date == "" {
		i.Date = time.Now().Format("20060102")
	}

	i.Logger.Info("import bill start", zap.String("target_date", i.Date))

	dir := fetcherDir + i.Date[0:6] + "/" + i.Date + "_inv.csv"
	rrecs, err := i.FileLoader.LoadRemixElectBillCSV(ctx, dir)
	if err != nil {
		i.Logger.Error("failed to load remix CSV", zap.Error(err))
		return err
	}

	i.Logger.Info("load remix bill CSV complete", zap.String("target_date", i.Date))

	for _, rrec := range rrecs {
		rec, err := rrec.ConvDBModel()
		if err != nil {
			i.Logger.Error("failed to convert to DB model", zap.Error(err))
			return err
		}
		err = i.DBRepository.AddElectBill(rec)
		if err != nil {
			i.Logger.Error("failed to insert record to DB", zap.Error(err))
			return err
		}
	}

	i.Logger.Info("DB insert bill complete")
	i.Logger.Info("import bill complete")
	return nil
}
