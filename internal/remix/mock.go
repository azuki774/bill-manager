package remix

import (
	"azuki774/bill-manager/internal/model"
	"context"
)

type mockDBRepository struct {
	Err error
}

type mockFileLoader struct {
	Err error
}

func (m *mockDBRepository) AddElectConsumption(record model.ElectConsumption) (err error) {
	if m.Err != nil {
		return m.Err
	}
	return nil
}

func (m *mockDBRepository) AddElectBill(record model.BillElect) (err error) {
	if m.Err != nil {
		return m.Err
	}
	return nil
}

func (m *mockFileLoader) LoadRemixElectConsumptionCSV(ctx context.Context, filePath string) (recs []model.RemixCSV, err error) {
	if m.Err != nil {
		return recs, m.Err
	}
	recs = []model.RemixCSV{
		{
			RecordDate:       "2022/01/01",
			TotalConsumption: 6, // kWh
			DayConsumption:   4, // kWh
			NightConsumption: 2, // kWh
		},
		{
			RecordDate:       "2022/01/02",
			TotalConsumption: 6, // kWh
			DayConsumption:   2, // kWh
			NightConsumption: 4, // kWh
		},
	}
	return recs, nil
}

func (m *mockFileLoader) LoadRemixElectBillCSV(ctx context.Context, filePath string) (recs []model.RemixBillingCSV, err error) {
	if m.Err != nil {
		return recs, m.Err
	}
	recs = []model.RemixBillingCSV{
		{
			BillingMonth:       "2022年12月分",
			ContractNumber:     "PP21XX",
			ProvidePointNumber: "0300XX",
			FacilityName:       "ばしょ",
			TotalConsumption:   202,
			Price:              "5,123",
		},
		{
			BillingMonth:       "2022年11月分",
			ContractNumber:     "PP21XX",
			ProvidePointNumber: "0300XX",
			FacilityName:       "ばしょ2",
			TotalConsumption:   202,
			Price:              "12,345",
		},
	}
	return recs, nil
}
