package api

import (
	"time"

	db "github.com/azuki774/bill-manager/internal/repository"
	"go.uber.org/zap"
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

	count, err := apis.remixdbR.GetCountElectConsume(tx, record.RecordDate)
	logger.Debug("count", count)
	if count > 0 {
		logger.Warnw("the data is already exists")
		err = db.ErrRecordAlreadyExists
		return err
	}
	if err != nil {
		logger.Error("error", zap.Error(err))
		return err
	}

	// add record
	err = apis.remixdbR.PostElectConsume(tx, record)
	if err != nil {
		logger.Errorw("post PostElectConsume data error", "error", err)
		return err
	}

	logger.Infow("post data to database", "data", record)
	return nil
}
