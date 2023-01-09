package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type GasBillingCSV struct {
	UsageMonthText string // 2022年12月分ガス料金詳細
	Price          string // 料金: "3,951"
	Consumption    string // 使用量: 123
}

func (g *GasBillingCSV) NewGasDBModel() (BillGas, error) {
	g.UsageMonthText = strings.Replace(g.UsageMonthText, "月分ガス料金詳細", "", -1)
	g.UsageMonthText = strings.Replace(g.UsageMonthText, "年", "", -1)
	g.Price = strings.Replace(g.Price, `"`, "", -1)
	g.Price = strings.Replace(g.Price, ",", "", -1)

	if len(g.UsageMonthText) == 5 {
		// 20231 -> 202301
		g.UsageMonthText = g.UsageMonthText[0:4] + "0" + g.UsageMonthText[4:5]
	}

	t, err := time.Parse("200601", g.UsageMonthText)
	if err != nil {
		return BillGas{}, err
	}
	t = t.AddDate(0, 1, 0) // 請求月 = 利用月 + 1

	price, err := strconv.Atoi(g.Price)
	if err != nil {
		return BillGas{}, err
	}

	cons, err := strconv.Atoi(g.Consumption)
	if err != nil {
		return BillGas{}, err
	}

	return BillGas{
		BillingMonth: t.Format("200601"),
		Price:        int64(price),
		Consumption:  int64(cons),
	}, nil
}

func NewGasBillingCSV(row []string) (rec GasBillingCSV, err error) {
	if len(row) != 4 {
		return GasBillingCSV{}, fmt.Errorf("invalid data")
	}

	return GasBillingCSV{
		UsageMonthText: row[0],
		Price:          row[1],
		Consumption:    row[2],
		// 支払い状況は取り込まない
	}, nil
}
