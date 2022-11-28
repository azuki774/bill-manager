package model

import (
	"context"
	"time"
)

type ShowRecord struct {
	Id           int64     `json:"id"`
	CategoryID   int64     `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Date         time.Time `json:"date"`
	Price        int64     `json:"price"`
	Memo         string    `json:"memo"`
}

// InCreateRecord is used to post to mawinter api
type CreateRecord struct {
	CategoryID int64  `json:"category_id"`
	Date       string `json:"date"` // YYYYMMDD
	Price      int64  `json:"price"`
	From       string `json:"from"`
	Type       string `json:"type"`
	Memo       string `json:"memo"`
}

// InCreateRecord is used for input JSON file struct
type InCreateRecord struct {
	CategoryID int64  `json:"category_id"`
	Day        string `json:"day"` // DD
	Price      int64  `json:"price"`
	Type       string `json:"type"`
	Memo       string `json:"memo"`
}

func (c *CreateRecord) FromInCreateRecord(ctx context.Context, inc *InCreateRecord) {
	c.CategoryID = inc.CategoryID
	c.Price = inc.Price
	c.Type = inc.Type
	c.From = "bill-manager-mawinter"
	c.Memo = inc.Memo
	if inc.Day != "" {
		c.Date = GetYYYYMM(ctx) + inc.Day
	}
}
