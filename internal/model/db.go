package model

import "time"

func (ElectConsumption) TableName() string {
	return "elect_consumption"
}

func (BillElect) TableName() string {
	return "bili_elect"
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
