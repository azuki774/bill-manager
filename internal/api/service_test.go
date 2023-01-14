package api

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"reflect"
	"testing"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var l *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	// config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	l, _ = config.Build()

	l.WithOptions(zap.AddStacktrace(zap.ErrorLevel))
}

func TestAPIService_GetBills(t *testing.T) {
	type fields struct {
		Logger *zap.Logger
		DBRepo DBRepository
	}
	type args struct {
		ctx    context.Context
		yyyymm string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantBills []model.BillAPIResponse
		wantErr   bool
	}{
		{
			name: "normal",
			fields: fields{
				Logger: l,
				DBRepo: &mockDBRepo{},
			},
			args: args{
				ctx:    context.Background(),
				yyyymm: "202201",
			},
			wantBills: []model.BillAPIResponse{
				{
					BillName: "elect",
					Price:    12345,
				},
				{
					BillName: "water",
					Price:    1234,
				},
				{
					BillName: "gas",
					Price:    123,
				},
			},
			wantErr: false,
		},
		{
			name: "not found",
			fields: fields{
				Logger: l,
				DBRepo: &mockDBRepo{err: model.ErrNotFound},
			},
			args: args{
				ctx:    context.Background(),
				yyyymm: "202201",
			},
			wantBills: []model.BillAPIResponse{},
			wantErr:   false,
		},
		{
			name: "internal error",
			fields: fields{
				Logger: l,
				DBRepo: &mockDBRepo{err: gorm.ErrInvalidData},
			},
			args: args{
				ctx:    context.Background(),
				yyyymm: "202201",
			},
			wantBills: []model.BillAPIResponse{},
			wantErr:   true,
		},
		{
			name: "invalid args",
			fields: fields{
				Logger: l,
				DBRepo: &mockDBRepo{},
			},
			args: args{
				ctx:    context.Background(),
				yyyymm: "2022",
			},
			wantBills: []model.BillAPIResponse{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := &APIService{
				Logger: tt.fields.Logger,
				DBRepo: tt.fields.DBRepo,
			}
			gotBills, err := ap.GetBills(tt.args.ctx, tt.args.yyyymm)
			if (err != nil) != tt.wantErr {
				t.Errorf("APIService.GetBills() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBills, tt.wantBills) {
				t.Errorf("APIService.GetBills() = %v, want %v", gotBills, tt.wantBills)
			}
		})
	}
}
