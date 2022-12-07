package remix

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

func TestImporter_Start(t *testing.T) {
	type fields struct {
		Logger       *zap.Logger
		DBRepository DBRepository
		FileLoader   FileLoader
		Date         string
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
			name: "ok",
			fields: fields{
				Logger:       l,
				DBRepository: &mockDBRepository{},
				FileLoader:   &mockFileLoader{},
				Date:         "",
			},
			args:    args{ctx: context.Background()},
			wantErr: false,
		},
		{
			name: "DB error",
			fields: fields{
				Logger:       l,
				DBRepository: &mockDBRepository{Err: fmt.Errorf("error")},
				FileLoader:   &mockFileLoader{},
				Date:         "",
			},
			args:    args{ctx: context.Background()},
			wantErr: true,
		},
		{
			name: "file load error",
			fields: fields{
				Logger:       l,
				DBRepository: &mockDBRepository{},
				FileLoader:   &mockFileLoader{Err: fmt.Errorf("error")},
				Date:         "",
			},
			args:    args{ctx: context.Background()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Importer{
				Logger:       tt.fields.Logger,
				DBRepository: tt.fields.DBRepository,
				FileLoader:   tt.fields.FileLoader,
				Date:         tt.fields.Date,
			}
			if err := i.Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Importer.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
