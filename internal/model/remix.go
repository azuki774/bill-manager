package model

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidData = errors.New("invalid data")
var ErrNotProvided = errors.New("not provided")

type RemixCSV struct {
	RecordDate       string // YYYY-MM-DD
	TotalConsumption int    // kWh
	DayConsumption   int    // kWh
	NightConsumption int    // kWh
}

type RemixBillingCSV struct {
	BillingMonth       string // 請求年月 20xx年yy月分
	ContractNumber     string // 契約番号
	ProvidePointNumber string // 供給地点特定番号
	FacilityName       string // 施設名称
	TotalConsumption   int    // 使用量(kWh)
	Price              string // 請求金額(円) 5,678 （カンマ区切り）
	// 請求書PDF
	// 電力使用CSV
}

func NewRemixCSV(row []string) (r RemixCSV, err error) {
	// 未計測(-) の値は ErrNotProvided を返す
	if len(row) != 4 {
		return RemixCSV{}, ErrInvalidData
	}
	r.RecordDate = row[0]

	if row[1] == "-" {
		return RemixCSV{}, ErrNotProvided
	}

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

func NewRemixBillingCSV(row []string) (r RemixBillingCSV, err error) {
	if len(row) != 8 {
		return RemixBillingCSV{}, ErrInvalidData
	}

	r.BillingMonth = row[0]
	r.ContractNumber = row[1]
	r.ProvidePointNumber = row[2]
	r.FacilityName = row[3]
	r.TotalConsumption, err = strconv.Atoi(row[4])
	if err != nil {
		return RemixBillingCSV{}, ErrInvalidData
	}
	r.Price = row[5]

	return r, nil
}

func (r *RemixBillingCSV) ConvDBModel() (record BillElect, err error) {
	bm := strings.ReplaceAll(r.BillingMonth, "年", "")
	bm = strings.ReplaceAll(bm, "月分", "")
	record.BillingMonth = bm                                   // 2022年11月分 -> 202211
	record.TotalConsumption = int64(r.TotalConsumption * 1000) // kWh -> Wh
	bp := strings.ReplaceAll(r.Price, ",", "")
	rp, err := strconv.Atoi(bp)
	if err != nil {
		return BillElect{}, ErrInvalidData
	}
	record.Price = int64(rp)
	return record, nil
}
