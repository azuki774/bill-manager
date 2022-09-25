package model

import "time"

type ShowRecord struct {
	Id           int64     `json:"id"`
	CategoryID   int64     `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Date         time.Time `json:"date"`
	Price        int64     `json:"price"`
	Memo         string    `json:"memo"`
}

type CreateRecord struct {
	CategoryID int64  `json:"category_id"`
	Date       string `json:"date"` // YYYYMMDD
	Price      int64  `json:"price"`
	Memo       string `json:"memo"`
}
