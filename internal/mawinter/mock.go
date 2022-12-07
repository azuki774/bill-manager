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
