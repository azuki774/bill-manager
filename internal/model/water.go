package model

type WaterBillingCSV struct {
	BillingMonth     string // 請求年月：4年 12月～5年 1月
	Price            string // 料金
	UsageTerm        string // 利用期間
	Consumption      string // 使用量
	DetailWaterPrice string // 内訳料金（水道）
	DetailSewerPrice string // 内訳料金（下水道）
}

func NewWaterDBModel(wbCSV WaterBillingCSV) (BillWater, error) {
	// TODO
	return BillWater{}, nil
}
