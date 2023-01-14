package model

import (
	"errors"
	"time"
)

var ErrInvalidData = errors.New("invalid data")
var ErrNotProvided = errors.New("not provided")
var ErrNotFound = errors.New("record not found")

func (ElectConsumption) TableName() string {
	return "elect_consumption"
}

func (BillElect) TableName() string {
	return "bill_elect"
}

func (BillWater) TableName() string {
	return "bill_water"
}

func (BillGas) TableName() string {
	return "bill_gas"
}

type ElectConsumption struct {
	ID               int64     `gorm:"column:id"`
	RecordDate       string    `gorm:"column:record_date"` // YYYYMMDD
	TotalConsumption int64     `gorm:"column:total_consumption"`
	DayConsumption   int64     `gorm:"column:day_consumption"`
	NightConsumption int64     `gorm:"column:night_consumption"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

type BillElect struct {
	ID               int64     `gorm:"column:id"`
	BillingMonth     string    `gorm:"column:billing_month"` // YYYYMMDD
	Price            int64     `gorm:"column:price"`
	TotalConsumption int64     `gorm:"column:total_consumption"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

type BillWater struct {
	ID               int64     `gorm:"column:id"`
	BillingMonth     string    `gorm:"column:billing_month"` // YYYYMMDD
	Price            int64     `gorm:"column:price"`
	Consumption      int64     `gorm:"column:consumption"`
	DetailWaterPrice int64     `gorm:"column:detail_water_price"`
	DetailSewerPrice int64     `gorm:"column:detail_sewer_price"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

type BillGas struct {
	ID           int64     `gorm:"column:id"`
	BillingMonth string    `gorm:"column:billing_month"` // YYYYMMDD
	Price        int64     `gorm:"column:price"`
	Consumption  int64     `gorm:"column:consumption"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
