package api

import (
	"errors"
	"time"

	db "github.com/azuki774/bill-manager/internal/db-ope"
	"gorm.io/gorm"
)

type RemixapiServiceRepository interface {
	GetElectConsume(date time.Time) (record db.ElectConsume, err error)
	PostElectConsume(record db.ElectConsume) (err error)
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

func (*UnimplementedremixapiService) PostElectConsume(record db.ElectConsume) (err error) {
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

func (apis *RemixapiServiceRepo) PostElectConsume(record db.ElectConsume) (err error) {
	tx := apis.remixdbR.OpenTx()
	defer apis.remixdbR.CloseTx(tx, err)

	_, err = apis.remixdbR.GetElectConsume(tx, record.RecordDate)
	if err != nil && (!errors.Is(err, gorm.ErrRecordNotFound)) {
		logger.Errorw("database unknown error", "error", err)
		return err
	}

	if err == nil {
		logger.Warnw("the record is already existed")
		return db.ErrRecordAlreadyExists
	}
	// TODO: 更新が必要かどうかを確認 and 更新 (not yet implemented)

	// add record
	err = apis.remixdbR.PostElectConsume(tx, record)
	if err != nil {
		logger.Errorw("post PostElectConsume data error", "error", err)
		return err
	}

	logger.Infow("post data to database", "data", record)
	return nil
}
