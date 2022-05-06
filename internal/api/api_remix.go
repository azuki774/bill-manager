package api

import (
	"time"

	db "github.com/azuki774/bill-manager/internal/db-ope"
)

type RemixapiService interface {
	GetElectConsume(date time.Time) (record db.ElectConsume, err error)
	PostElectConsume(date time.Time, record db.ElectConsume) (err error)
	mustEmbedUnimplementedElectConsumeService()
}
type remixapiService struct {
	remixdbR db.ElectConsumeDBRepository
	UnimplementedremixapiService
}

type UnimplementedremixapiService struct {
}

func (*UnimplementedremixapiService) mustEmbedUnimplementedElectConsumeService() {}

func (*UnimplementedremixapiService) GetElectConsume(date time.Time) (record db.ElectConsume, err error) {
	return db.ElectConsume{}, nil
}

func (*UnimplementedremixapiService) PostElectConsume(date time.Time, record db.ElectConsume) (err error) {
	return nil
}

type UnsafeElectConsumeService interface {
	mustEmbedUnimplementedElectConsumeService()
}

func NewRemixapiService(dbR_in db.ElectConsumeDBRepository) RemixapiService {
	return &remixapiService{remixdbR: dbR_in}
}

func (apis *remixapiService) GetElectConsume(date time.Time) (record db.ElectConsume, err error) {
	tx := apis.remixdbR.OpenTx()
	defer apis.remixdbR.CloseTx(tx, err)

	record, err = apis.remixdbR.GetElectConsume(tx, date)
	if err != nil {
		// TODO: エラー内容で分ける
		logger.Error("error", err)
		return record, err
	}
	return record, nil
}

func (apis *remixapiService) PostElectConsume(date time.Time, record db.ElectConsume) (err error) {
	return nil
}
