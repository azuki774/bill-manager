package water

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/zap"
)

var l *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	// config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	l, _ = config.Build()

	l.WithOptions(zap.AddStacktrace(zap.ErrorLevel))
}
func TestWaterService_Import(t *testing.T) {
	type fields struct {
		Logger        *zap.Logger
		DBRepository  DBRepository
		FileLoader    FileLoader
		Date          string
		Downloader    Downloader
		remoteRootDir string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Logger:       l,
				DBRepository: &mockDBRepository{},
				FileLoader:   &mockFileLoader{},
				Date:         "20230101",
				Downloader:   &mockDownloader{},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "DB error",
			fields: fields{
				Logger:       l,
				DBRepository: &mockDBRepository{err: fmt.Errorf("error")},
				FileLoader:   &mockFileLoader{},
				Date:         "20230101",
				Downloader:   &mockDownloader{},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "data error",
			fields: fields{
				Logger:       l,
				DBRepository: &mockDBRepository{},
				FileLoader:   &mockFileLoader{err: fmt.Errorf("error")},
				Date:         "20230101",
				Downloader:   &mockDownloader{},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "download error",
			fields: fields{
				Logger:       l,
				DBRepository: &mockDBRepository{},
				FileLoader:   &mockFileLoader{},
				Date:         "20230101",
				Downloader:   &mockDownloader{err: fmt.Errorf("error")},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WaterService{
				Logger:        tt.fields.Logger,
				DBRepository:  tt.fields.DBRepository,
				FileLoader:    tt.fields.FileLoader,
				Date:          tt.fields.Date,
				Downloader:    tt.fields.Downloader,
				remoteRootDir: tt.fields.remoteRootDir,
			}
			if err := w.Import(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("WaterService.Import() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
