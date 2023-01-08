package water

import (
	"context"

	"go.uber.org/zap"
)

type DBRepository interface {
}
type FileLoader interface {
}
type Downloader interface {
}

type WaterService struct {
	Logger       *zap.Logger
	DBRepository DBRepository
	FileLoader   FileLoader
	Date         string // YYYYMMDD
	Downloader   Downloader
}

func (w *WaterService) Import(ctx context.Context) (err error) {
	return nil
}
