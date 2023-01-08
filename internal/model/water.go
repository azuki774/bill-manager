package model

import (
	"strconv"
	"strings"
)

type WaterBillingCSV struct {
	BillingMonth     string // 請求年月：4年12月 ～ 5年1月分
	Price            string // 料金： "3,951"
	UsageTerm        string // 利用期間：,11月 2日 ～  1月 4日 (64日間)
	Consumption      string // 使用量：16
	DetailWaterPrice string // 内訳料金（水道）："2,719"
	DetailSewerPrice string // 内訳料金（下水道）："2,719"
}

func (w *WaterBillingCSV) NewWaterDBModel() (BillWater, error) {
	var bw BillWater
	w.Price = strings.Replace(w.Price, ",", "", -1)
	w.Price = strings.Replace(w.Price, `"`, "", -1)
	w.DetailWaterPrice = strings.Replace(w.DetailWaterPrice, ",", "", -1)
	w.DetailWaterPrice = strings.Replace(w.DetailWaterPrice, `"`, "", -1)
	w.DetailSewerPrice = strings.Replace(w.DetailSewerPrice, ",", "", -1)
	w.DetailSewerPrice = strings.Replace(w.DetailSewerPrice, `"`, "", -1)

	// bmList := strings.Split(w.BillingMonth, " ") // 4年12月 ～ 5年1月分 -> [4年12月,～,5年1月分]
	// bmStr = bmList[len(bmList)-1]                // [4年12月,～,5年1月分] -> 5年1月分

	price, err := strconv.Atoi(w.Price)
	if err != nil {
		return BillWater{}, err
	}

	cons, err := strconv.Atoi(w.Consumption)
	if err != nil {
		return BillWater{}, err
	}

	detailWater, err := strconv.Atoi(w.DetailWaterPrice)
	if err != nil {
		return BillWater{}, err
	}

	detailSewer, err := strconv.Atoi(w.DetailSewerPrice)
	if err != nil {
		return BillWater{}, err
	}

	// bw.BillingMonth =
	bw.Price = int64(price)
	bw.Consumption = int64(cons)
	bw.DetailWaterPrice = int64(detailWater)
	bw.DetailSewerPrice = int64(detailSewer)
	return bw, nil
}
