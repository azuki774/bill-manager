package model

import (
	"reflect"
	"testing"
)

func TestGasBillingCSV_NewGasDBModel(t *testing.T) {
	type fields struct {
		UsageMonthText string
		Price          string
		Consumption    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    BillGas
		wantErr bool
	}{
		{
			name: "2022年12月分ガス料金詳細",
			fields: fields{
				UsageMonthText: "2022年12月分ガス料金詳細",
				Price:          "1,234",
				Consumption:    "123",
			},
			want: BillGas{
				BillingMonth: "202301",
				Price:        1234,
				Consumption:  123,
			},
			wantErr: false,
		},
		{
			name: "2023年1月分ガス料金詳細",
			fields: fields{
				UsageMonthText: "2023年1月分ガス料金詳細",
				Price:          "1,234",
				Consumption:    "123",
			},
			want: BillGas{
				BillingMonth: "202302",
				Price:        1234,
				Consumption:  123,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GasBillingCSV{
				UsageMonthText: tt.fields.UsageMonthText,
				Price:          tt.fields.Price,
				Consumption:    tt.fields.Consumption,
			}
			got, err := g.NewGasDBModel()
			if (err != nil) != tt.wantErr {
				t.Errorf("GasBillingCSV.NewGasDBModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GasBillingCSV.NewGasDBModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
