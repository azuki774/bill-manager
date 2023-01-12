package api

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"errors"

	"go.uber.org/zap"
)

type DBRepository interface {
	GetBillElect(yyyymm string) (b model.BillElect, err error)
	GetBillGas(yyyymm string) (b model.BillGas, err error)
	GetBillWater(yyyymm string) (b model.BillWater, err error)
}

type APIService struct {
	Logger *zap.Logger
	DBRepo DBRepository
}

func (ap *APIService) GetBills(ctx context.Context, yyyymm string) (bills []model.BillAPIResponse, err error) {
	ap.Logger.Info("get bill start")
	bills = []model.BillAPIResponse{}

	if err := model.ValidYYYYMM(yyyymm); err != nil {
		ap.Logger.Error("invalid args", zap.String("yyyymm", yyyymm), zap.Error(err))
		return []model.BillAPIResponse{}, err
	}

	// Get Bill Elect
	b, err := ap.DBRepo.GetBillElect(yyyymm)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		ap.Logger.Error("failed to get bill elect from DB", zap.Error(err))
		return []model.BillAPIResponse{}, err
	} else if err == nil {
		// record found
		bills = append(bills, b.NewBillAPIResponse())
	}

	// Get Bill Water
	w, err := ap.DBRepo.GetBillWater(yyyymm)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		ap.Logger.Error("failed to get bill water from DB", zap.Error(err))
		return []model.BillAPIResponse{}, err
	} else if err == nil {
		// record found
		bills = append(bills, w.NewBillAPIResponse())
	}
	// Get Bill Gas
	g, err := ap.DBRepo.GetBillGas(yyyymm)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		ap.Logger.Error("failed to get bill gas from DB", zap.Error(err))
		return []model.BillAPIResponse{}, err
	} else if err == nil {
		// record found
		bills = append(bills, g.NewBillAPIResponse())
	}

	ap.Logger.Info("get bill records successfully")
	return bills, nil
}
