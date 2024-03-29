package factory

import (
	"fmt"
	"os"

	"azuki774/bill-manager/internal/api"
	"azuki774/bill-manager/internal/downloader"
	"azuki774/bill-manager/internal/gas"
	"azuki774/bill-manager/internal/mawinter"
	"azuki774/bill-manager/internal/remix"
	"azuki774/bill-manager/internal/repository"
	"azuki774/bill-manager/internal/server"
	"azuki774/bill-manager/internal/water"

	"go.uber.org/zap"
)

func NewUsecaseMawinter(h *repository.HTTPClient, f *repository.FileLoader) (u *mawinter.UsecaseMawinter, err error) {
	l, err := NewLogger()
	if err != nil {
		fmt.Printf("failed to get logger: %v\n", err)
		return nil, err
	}

	return &mawinter.UsecaseMawinter{Logger: l, HTTPClient: h, FileLoader: f}, nil
}

func NewFileLoader() *repository.FileLoader {
	return &repository.FileLoader{}
}

func NewSFTPDownloader() *downloader.SFTPClient {
	return &downloader.SFTPClient{Host: os.Getenv("SRC_HOST")}
}

func NewUsecaseRemix(l *zap.Logger, d *repository.DBRepository, f *repository.FileLoader) (r *remix.Importer) {
	return &remix.Importer{Logger: l, FileLoader: f, DBRepository: d}
}

func NewUsecaseWater(l *zap.Logger, d *repository.DBRepository, f *repository.FileLoader, date string) (u *water.WaterService) {
	return &water.WaterService{Logger: l, DBRepository: d, FileLoader: f, Date: date}
}

func NewUsecaseGas(l *zap.Logger, d *repository.DBRepository, f *repository.FileLoader, date string) (u *gas.GasService) {
	return &gas.GasService{Logger: l, DBRepository: d, FileLoader: f, Date: date}
}

func NewAPIService(l *zap.Logger, d *repository.DBRepository) (ap *api.APIService) {
	return &api.APIService{Logger: l, DBRepo: d}
}

func NewAPIServer(l *zap.Logger, port string, ap *api.APIService) *server.Server { // TODO
	return &server.Server{Logger: l, Port: port, APISvc: ap}
}
