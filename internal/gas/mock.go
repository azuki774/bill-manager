package remix

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

func (m *mockDBRepository) AddGasBill(r model.BillGas) (err error) {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m *mockFileLoader) LoadGasBillCSV(ctx context.Context, dir string) (recs []model.GasBillingCSV, err error) {
	if m.err != nil {
		return []model.GasBillingCSV{}, m.err
	}

	recs = []model.GasBillingCSV{
		{
			UsageMonthText: "2022年12月分ガス料金詳細",
			Price:          `"1,234"`,
			Consumption:    "16",
		},
	}
	return recs, nil
}
