package model

import (
	"reflect"
	"testing"
)

func TestWaterBillingCSV_NewWaterDBModel(t *testing.T) {
	type fields struct {
		BillingMonth     string
		Price            string
		UsageTerm        string
		Consumption      string
		DetailWaterPrice string
		DetailSewerPrice string
	}
	tests := []struct {
		name    string
		fields  fields
		want    BillWater
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WaterBillingCSV{
				BillingMonth:     tt.fields.BillingMonth,
				Price:            tt.fields.Price,
				UsageTerm:        tt.fields.UsageTerm,
				Consumption:      tt.fields.Consumption,
				DetailWaterPrice: tt.fields.DetailWaterPrice,
				DetailSewerPrice: tt.fields.DetailSewerPrice,
			}
			got, err := w.NewWaterDBModel()
			if (err != nil) != tt.wantErr {
				t.Errorf("WaterBillingCSV.NewWaterDBModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WaterBillingCSV.NewWaterDBModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
