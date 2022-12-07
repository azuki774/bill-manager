package model

import (
	"errors"
	"strconv"
	"time"
)

var ErrInvalidData = errors.New("invalid data")

type RemixCSV struct {
	RecordDate       string // YYYY-MM-DD
	TotalConsumption int    // kWh
	DayConsumption   int    // kWh
	NightConsumption int    // kWh
}

func NewRemixCSV(row []string) (r RemixCSV, err error) {
	if len(row) != 4 {
		return RemixCSV{}, ErrInvalidData
	}
	r.RecordDate = row[0]
	r.TotalConsumption, err = strconv.Atoi(row[1])
	if err != nil {
		return RemixCSV{}, ErrInvalidData
	}

	r.DayConsumption, err = strconv.Atoi(row[2])
	if err != nil {
		return RemixCSV{}, ErrInvalidData
	}

	r.NightConsumption, err = strconv.Atoi(row[3])
	if err != nil {
		return RemixCSV{}, ErrInvalidData
	}

	return r, nil
}

func (r *RemixCSV) ConvDBModel() (record ElectConsumption, err error) {
	rd, err := time.Parse("2006/01/02", r.RecordDate)
	if err != nil {
		return ElectConsumption{}, err
	}
	record.RecordDate = rd.Format("20060102")
	record.TotalConsumption = int64(r.TotalConsumption * 1000) // kWh -> Wh
	record.DayConsumption = int64(r.DayConsumption * 1000)     // kWh -> Wh
	record.NightConsumption = int64(r.NightConsumption * 1000) // kWh -> Wh
	return record, nil
}
