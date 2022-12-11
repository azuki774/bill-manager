package mawinter

import (
	"context"

	"azuki774/bill-manager/internal/model"
)

type MockHTTPClient struct {
	resBody    []byte
	statusCode int
	err        error
}

func (m *MockHTTPClient) PostJson(ctx context.Context, endPoint string, reqBody []byte) (resBody []byte, statusCode int, err error) {
	return m.resBody, m.statusCode, m.err
}

type MockFileLoader struct {
	recs []model.CreateRecord
	err  error
}

func (m MockFileLoader) LoadRecordsFromJSON(ctx context.Context, filePath string) (recs []model.CreateRecord, err error) {
	return m.recs, m.err
}

type MockDBRepository struct {
	err error
}

func (m MockDBRepository) GetElectBillFromDB(ctx context.Context, billingMonth string) (price int, err error) {
	if m.err != nil {
		return 0, m.err
	}
	return 1000, nil
}
