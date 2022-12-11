package mawinter

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"
)

var l *zap.Logger

func TestMain(m *testing.M) {
	config := zap.NewProductionConfig()
	// config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	l, _ = config.Build()

	l.WithOptions(zap.AddStacktrace(zap.ErrorLevel))

	m.Run()
}

func TestUsecaseMawinter_RegistFromJSON(t *testing.T) {
	type fields struct {
		Logger     *zap.Logger
		HTTPClient HTTPClient
		FileLoader FileLoader
	}
	type args struct {
		ctx      context.Context
		jsonfile string
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
				Logger: l,
				HTTPClient: &MockHTTPClient{
					resBody:    []byte("test"),
					statusCode: http.StatusCreated,
					err:        nil,
				},
				FileLoader: MockFileLoader{
					recs: []model.CreateRecord{
						{
							CategoryID: 100,
							Price:      100,
							From:       "bill-manager-mawinter",
							Memo:       "Memo",
						},
						{
							CategoryID: 200,
							Price:      200,
							From:       "bill-manager-mawinter",
							Memo:       "Memo",
						},
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				jsonfile: "testfile1",
			},
			wantErr: false,
		},
		{
			name: "not authorized",
			fields: fields{
				Logger: l,
				HTTPClient: &MockHTTPClient{
					resBody:    []byte("test"),
					statusCode: 400,
					err:        nil,
				},
				FileLoader: MockFileLoader{
					recs: []model.CreateRecord{
						{
							CategoryID: 100,
							Price:      100,
							From:       "bill-manager-mawinter",
							Memo:       "Memo",
						},
						{
							CategoryID: 200,
							Price:      200,
							From:       "bill-manager-mawinter",
							Memo:       "Memo",
						},
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				jsonfile: "testfile1",
			},
			wantErr: true,
		},
		{
			name: "failed post",
			fields: fields{
				Logger: l,
				HTTPClient: &MockHTTPClient{
					err: errors.New("error"),
				},
				FileLoader: MockFileLoader{
					recs: []model.CreateRecord{
						{
							CategoryID: 100,
							Price:      100,
							From:       "bill-manager-mawinter",
							Memo:       "Memo",
						},
						{
							CategoryID: 200,
							Price:      200,
							From:       "bill-manager-mawinter",
							Memo:       "Memo",
						},
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				jsonfile: "testfile1",
			},
			wantErr: true,
		},
		{
			name: "failed load",
			fields: fields{
				Logger: l,
				HTTPClient: &MockHTTPClient{
					resBody:    []byte("test"),
					statusCode: 201,
					err:        nil,
				},
				FileLoader: MockFileLoader{
					err: errors.New("unknown error"),
				},
			},
			args: args{
				ctx:      context.Background(),
				jsonfile: "testfile1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsecaseMawinter{
				Logger:     tt.fields.Logger,
				HTTPClient: tt.fields.HTTPClient,
				FileLoader: tt.fields.FileLoader,
			}
			if err := u.RegistFromJSON(tt.args.ctx, tt.args.jsonfile); (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMawinter.RegistFromJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecaseMawinter_RegistElectBill(t *testing.T) {
	type fields struct {
		Logger       *zap.Logger
		HTTPClient   HTTPClient
		FileLoader   FileLoader
		DBRepository DBRepository
	}
	type args struct {
		ctx          context.Context
		billingMonth string
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
				DBRepository: &MockDBRepository{},
				HTTPClient: &MockHTTPClient{
					resBody:    []byte(""),
					statusCode: 201,
					err:        nil,
				},
			},
			args: args{
				ctx:          context.Background(),
				billingMonth: "202211",
			},
			wantErr: false,
		},
		{
			name: "unexpected http status code",
			fields: fields{
				Logger:       l,
				DBRepository: &MockDBRepository{},
				HTTPClient: &MockHTTPClient{
					resBody:    []byte(""),
					statusCode: 500,
					err:        nil,
				},
			},
			args: args{
				ctx:          context.Background(),
				billingMonth: "202211",
			},
			wantErr: true,
		},
		{
			name: "DB error",
			fields: fields{
				Logger: l,
				DBRepository: &MockDBRepository{
					err: fmt.Errorf("error"),
				},
				HTTPClient: &MockHTTPClient{
					resBody:    []byte(""),
					statusCode: 201,
					err:        nil,
				},
			},
			args: args{
				ctx:          context.Background(),
				billingMonth: "202211",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsecaseMawinter{
				Logger:       tt.fields.Logger,
				HTTPClient:   tt.fields.HTTPClient,
				FileLoader:   tt.fields.FileLoader,
				DBRepository: tt.fields.DBRepository,
			}
			if err := u.RegistElectBill(tt.args.ctx, tt.args.billingMonth); (err != nil) != tt.wantErr {
				t.Errorf("UsecaseMawinter.RegistElectBill() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
