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
		{
			name: "ok",
			fields: fields{
				BillingMonth:     "4年12月 ～ 5年1月分",
				Price:            `"1,234`,
				UsageTerm:        "11月 2日 ～  1月 4日 (64日間)",
				Consumption:      "10",
				DetailWaterPrice: `"1,234"`,
				DetailSewerPrice: `"1,234"`,
			},
			want: BillWater{
				BillingMonth:     "202301",
				Price:            1234,
				Consumption:      10,
				DetailWaterPrice: 1234,
				DetailSewerPrice: 1234,
			},
			wantErr: false,
		},
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

func Test_toBillMonthWareki(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "5年1月分",
			args: args{
				str: "5年1月分",
			},
			want:    "202301",
			wantErr: false,
		},
		{
			name: "6年12月分",
			args: args{
				str: "6年12月分",
			},
			want:    "202412",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBillMonthWareki(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("toBillMonthWareki() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toBillMonthWareki() = %v, want %v", got, tt.want)
			}
		})
	}
}
