package mawinter

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"azuki774/bill-manager/internal/model"

	"go.uber.org/zap"
)

type HTTPClient interface {
	PostJson(ctx context.Context, endPoint string, reqBody []byte) (resBody []byte, statusCode int, err error)
}

type FileLoader interface {
	LoadRecordsFromJSON(ctx context.Context, filePath string) (recs []model.CreateRecord, err error)
}

type DBRepository interface {
	GetElectBillFromDB(ctx context.Context, billingMonth string) (price int, err error)
}

type UsecaseMawinter struct {
	Logger       *zap.Logger
	HTTPClient   HTTPClient
	FileLoader   FileLoader
	DBRepository DBRepository
}

func (u *UsecaseMawinter) RegistFromJSON(ctx context.Context, jsonfile string) (err error) {
	u.Logger.Info("regist records from JSON")
	recs, err := u.FileLoader.LoadRecordsFromJSON(ctx, jsonfile)
	if err != nil {
		u.Logger.Error("failed to load jsonfile", zap.String("filename", jsonfile), zap.Error(err))
		return err
	}

	for i, rec := range recs {
		iLogger := u.Logger.With(zap.Int("index", i))
		b, err := json.Marshal(rec)
		if err != nil {
			iLogger.Error("failed to marshal JSON", zap.Error(err))
			return err
		}

		resBody, statusCode, err := u.HTTPClient.PostJson(ctx, "/record/", b)
		if err != nil {
			iLogger.Error("failed to post records", zap.Error(err))
			return err
		}
		if statusCode != http.StatusCreated {
			err = errors.New("unexpected status code")
			iLogger.Error("unexpected status code", zap.Int("status_code", statusCode), zap.Error(err))
			return err
		}
		iLogger.Info("post records", zap.String("body", string(resBody)))
	}

	return nil
}

func (u *UsecaseMawinter) RegistElectBill(ctx context.Context, billingMonth string) (err error) {
	if billingMonth == "" {
		// デフォルト１ヶ月前
		t := time.Now()
		billingMonth = t.AddDate(0, 0, -1).Format("200601")
	}

	price, err := u.DBRepository.GetElectBillFromDB(ctx, billingMonth)
	if err != nil {
		u.Logger.Error("failed to get bill from DB", zap.Error(err))
		return err
	}

	rec := model.NewElectCreateRecord(price)
	b, err := json.Marshal(rec)
	if err != nil {
		u.Logger.Error("failed to marshal JSON", zap.Error(err))
		return err
	}

	resBody, statusCode, err := u.HTTPClient.PostJson(ctx, "/record/", b)
	if err != nil {
		u.Logger.Error("failed to post records", zap.Error(err))
		return err
	}
	if statusCode != http.StatusCreated {
		err = errors.New("unexpected status code")
		u.Logger.Error("unexpected status code", zap.Int("status_code", statusCode), zap.Error(err))
		return err
	}

	u.Logger.Info("post records", zap.String("body", string(resBody)))
	return nil
}
