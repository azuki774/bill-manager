package api

import (
	"azuki774/bill-manager/internal/model"
	"time"
)

type mockDBRepo struct {
	err error
}

func (m *mockDBRepo) GetBillElect(yyyymm string) (b model.BillElect, err error) {
	if m.err != nil {
		return model.BillElect{}, m.err
	}
	return model.BillElect{
		ID:               1,
		BillingMonth:     "202201", // YYYYMMDD
		Price:            12345,
		TotalConsumption: 100,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}, nil
}

func (m *mockDBRepo) GetBillWater(yyyymm string) (b model.BillWater, err error) {
	if m.err != nil {
		return model.BillWater{}, m.err
	}

	return model.BillWater{
		ID:           2,
		BillingMonth: "202201", // YYYYMMDD
		Price:        1234,
		Consumption:  100,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

func (m *mockDBRepo) GetBillGas(yyyymm string) (b model.BillGas, err error) {
	if m.err != nil {
		return model.BillGas{}, m.err
	}

	return model.BillGas{
		ID:           3,
		BillingMonth: "202201", // YYYYMMDD
		Price:        123,
		Consumption:  100,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
