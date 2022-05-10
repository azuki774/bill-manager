package api

import (
	"time"

	db "github.com/azuki774/bill-manager/internal/db-ope"
)

type RemixapiServiceRepository interface {
	GetElectConsume(date time.Time) (record db.ElectConsume, err error)
	PostElectConsume(date time.Time, record db.ElectConsume) (err error)
	mustEmbedUnimplementedElectConsumeService()
}
type RemixapiServiceRepo struct {
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

func NewRemixapiService(dbR_in db.ElectConsumeDBRepository) RemixapiServiceRepository {
	return &RemixapiServiceRepo{remixdbR: dbR_in}
}

func (apis *RemixapiServiceRepo) GetElectConsume(date time.Time) (record db.ElectConsume, err error) {
	tx := apis.remixdbR.OpenTx()
	defer apis.remixdbR.CloseTx(tx, err)

	record, err = apis.remixdbR.GetElectConsume(tx, date)
	if err != nil {
		// TODO: エラー内容で分ける
		logger.Errorw("fetch ElectConsume data error", "error", err)
		return db.ElectConsume{}, err
	}

	logger.Debugw("Get ElectConsume data from DB", "data", record)
	return record, nil
}

func (apis *RemixapiServiceRepo) PostElectConsume(date time.Time, record db.ElectConsume) (err error) {
	return nil
}
