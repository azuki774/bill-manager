package db_ope

import (
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
	GetElectConsume(tx *gorm.DB, t time.Time) (record ElectConsume, err error)
	PostElectConsume(tx *gorm.DB, record ElectConsume) (err error)
	mustEmbedUnimplementedElectConsumeDBRepository()
}

type electConsumeDBRepository struct {
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

func (dbR *UnimplementedElectConsumeDBRepository) mustEmbedUnimplementedElectConsumeDBRepository() {}

type UnsafeElectConsumeDBRepository interface {
	mustEmbedUnimplementedElectConsumeDBRepository()
}

func NewDBRepository(conn *gorm.DB) ElectConsumeDBRepository {
	return &electConsumeDBRepository{conn: conn}
}

func (ecdbR *electConsumeDBRepository) OpenTx() *gorm.DB {
	tx := ecdbR.conn.Begin()
	return tx
}

func (ecdbR *electConsumeDBRepository) CloseTx(tx *gorm.DB, err error) error {
	if err != nil {
		logger.Error("CloseTx (Rollback): %s", zap.Error(err))
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return nil
}
