package model

import "time"

type ElectConsumption struct {
	ID               int64     `gorm:"id,primaryKey"`
	RecordDate       string    `gorm:"column:record_date,index,not null,unique"` // YYYYMMDD
	TotalConsumption int64     `gorm:"column:total_consumption,not null"`
	DayConsumption   int64     `gorm:"column:day_consumption"`
	NightConsumption int64     `gorm:"column:night_consumption"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}
