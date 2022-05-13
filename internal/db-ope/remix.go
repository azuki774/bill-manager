package db_ope

import (
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ElectConsume struct {
	Id         int64 `gorm:"primaryKey"`
	RecordDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Daytime    int64
	Nighttime  int64
	Total      int64
}

type ElectConsumeDBRepository interface {
	OpenTx() *gorm.DB
	CloseTx(tx *gorm.DB, err error) error
	GetElectConsume(tx *gorm.DB, t time.Time) (record ElectConsume, err error) // SELECT * FROM `elect_consumes` WHERE record_date BETWEEN '2000-10-01 00:00:00' AND '2000-10-01 23:59:59';
	PostElectConsume(tx *gorm.DB, record ElectConsume) (err error)             // INSERT INTO `elect_consumes` (`Record_date`,`Daytime`,`Nighttime`,`Total`) VALUES ("2000-10-01", 111, 222, 333);
	GetCountElectConsume(tx *gorm.DB, t time.Time) (count int64, err error)    // SELECT COUNT(*) FROM `elect_consumes` WHERE record_date BETWEEN '2000-10-01 00:00:00' AND '2000-10-01 23:59:59';
	mustEmbedUnimplementedElectConsumeDBRepository()
}

type ElectConsumeDBRrepo struct {
	conn *gorm.DB
	UnimplementedElectConsumeDBRepository
}

type UnimplementedElectConsumeDBRepository struct {
}

func (*UnimplementedElectConsumeDBRepository) OpenTx() *gorm.DB {
	return nil
}

func (*UnimplementedElectConsumeDBRepository) CloseTx(tx *gorm.DB, err error) error {
	return nil
}

func (*UnimplementedElectConsumeDBRepository) GetElectConsume(tx *gorm.DB, t time.Time) (record ElectConsume, err error) {
	return ElectConsume{}, nil
}

func (*UnimplementedElectConsumeDBRepository) PostElectConsume(tx *gorm.DB, record ElectConsume) (err error) {
	return nil
}

func (*UnimplementedElectConsumeDBRepository) GetCountElectConsume(tx *gorm.DB, t time.Time) (count int64, err error) {
	return 0, nil
}

func (dbR *UnimplementedElectConsumeDBRepository) mustEmbedUnimplementedElectConsumeDBRepository() {}

type UnsafeElectConsumeDBRepository interface {
	mustEmbedUnimplementedElectConsumeDBRepository()
}

func NewElectConsumeDBRepository(conn *gorm.DB) ElectConsumeDBRepository {
	return &ElectConsumeDBRrepo{conn: conn}
}

func (dbR *ElectConsumeDBRrepo) OpenTx() *gorm.DB {
	tx := dbR.conn.Begin()
	return tx
}

func (dbR *ElectConsumeDBRrepo) CloseTx(tx *gorm.DB, err error) error {
	if err != nil {
		logger.Errorw("CloseTx (Rollback)", "error", zap.Error(err))
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return nil
}

func (dbR *ElectConsumeDBRrepo) GetElectConsume(tx *gorm.DB, t time.Time) (record ElectConsume, err error) {
	logger.Debugw("GetElectConsume")
	tFirst := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Now().Location())
	tEnd := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Now().Location())
	err = tx.Where("record_date BETWEEN ? AND ?", tFirst, tEnd).First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ElectConsume{}, gorm.ErrRecordNotFound
	} else if err != nil {
		logger.Errorw("database internal error", "error", err)
		return ElectConsume{}, err
	}
	return record, err
}

func (dbR *ElectConsumeDBRrepo) PostElectConsume(tx *gorm.DB, record ElectConsume) (err error) {
	logger.Debugw("PostElectConsume")
	res := tx.Create(&record)
	if res.Error != nil {
		logger.Errorw("database internal error", "error", err)
		return ErrInternal
	}
	return nil
}

func (dbR *ElectConsumeDBRrepo) GetCountElectConsume(tx *gorm.DB, t time.Time) (count int64, err error) {
	logger.Debugw("GetCountElectConsume")
	tFirst := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Now().Location())
	tEnd := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Now().Location())
	tx.Model(&ElectConsume{}).Where("record_date BETWEEN ? AND ?", tFirst, tEnd).Count(&count)
	return count, nil
}
