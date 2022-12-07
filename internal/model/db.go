package model

import "time"

func (ElectConsumption) TableName() string {
	return "elect_consumption"
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
