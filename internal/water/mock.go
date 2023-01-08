package water

import (
	"azuki774/bill-manager/internal/model"
	"context"
)

type mockDBRepository struct {
	err error
}
type mockFileLoader struct {
	err error
}
type mockDownloader struct {
	err error
}

func (m *mockDBRepository) AddWaterBill(r model.BillWater) (err error) {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m *mockFileLoader) LoadWaterBillCSV(ctx context.Context, dir string) (recs []model.WaterBillingCSV, err error) {
	if m.err != nil {
		return []model.WaterBillingCSV{}, m.err
	}

	recs = []model.WaterBillingCSV{
		{
			BillingMonth:     "4年12月 ～ 5年1月分",
			Price:            `"3,951"`,
			UsageTerm:        "11月 2日 ～  1月 4日 (64日間)",
			Consumption:      "16",
			DetailWaterPrice: `"2,719"`,
			DetailSewerPrice: `"2,719"`,
		},
	}
	return recs, nil
}

func (m *mockDownloader) Download(ctx context.Context, dir string, remoteDir string) (err error) {
	if m.err != nil {
		return m.err
	}
	return nil
}
