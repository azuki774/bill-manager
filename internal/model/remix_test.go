package model

import (
	"reflect"
	"testing"
)

func TestNewRemixCSV(t *testing.T) {
	type args struct {
		row []string
	}
	tests := []struct {
		name    string
		args    args
		wantR   RemixCSV
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				row: []string{
					"2022/12/01",
					"6",
					"4",
					"2",
				},
			},
			wantR: RemixCSV{
				RecordDate:       "2022/12/01",
				TotalConsumption: 6,
				DayConsumption:   4,
				NightConsumption: 2,
			},
			wantErr: false,
		},
		{
			name: "invalid args 1",
			args: args{
				row: []string{
					"2022/12/01",
					"6",
					"4",
					"3.2",
				},
			},
			wantR:   RemixCSV{},
			wantErr: true,
		},
		{
			name: "invalid args 2",
			args: args{
				row: []string{
					"2022/12/01",
					"6",
					"4",
				},
			},
			wantR:   RemixCSV{},
			wantErr: true,
		},
		{
			name: "not provided",
			args: args{
				row: []string{
					"2022/12/01",
					"-",
					"-",
					"-",
				},
			},
			wantR:   RemixCSV{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := NewRemixCSV(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRemixCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewRemixCSV() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestRemixCSV_ConvDBModel(t *testing.T) {
	type fields struct {
		RecordDate       string
		TotalConsumption int
		DayConsumption   int
		NightConsumption int
	}
	tests := []struct {
		name       string
		fields     fields
		wantRecord ElectConsumption
		wantErr    bool
	}{
		{
			name: "ok",
			fields: fields{
				RecordDate:       "2022/12/01",
				TotalConsumption: 6,
				DayConsumption:   4,
				NightConsumption: 2,
			},
			wantRecord: ElectConsumption{
				RecordDate:       "20221201",
				TotalConsumption: 6000,
				DayConsumption:   4000,
				NightConsumption: 2000,
			},
			wantErr: false,
		},
		{
			name: "invalid args 1",
			fields: fields{
				RecordDate:       "2022-12-01",
				TotalConsumption: 6,
				DayConsumption:   4,
				NightConsumption: 2,
			},
			wantRecord: ElectConsumption{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RemixCSV{
				RecordDate:       tt.fields.RecordDate,
				TotalConsumption: tt.fields.TotalConsumption,
				DayConsumption:   tt.fields.DayConsumption,
				NightConsumption: tt.fields.NightConsumption,
			}
			gotRecord, err := r.ConvDBModel()
			if (err != nil) != tt.wantErr {
				t.Errorf("RemixCSV.ConvDBModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("RemixCSV.ConvDBModel() = %v, want %v", gotRecord, tt.wantRecord)
			}
		})
	}
}

func TestNewRemixBillingCSV(t *testing.T) {
	type args struct {
		row []string
	}
	tests := []struct {
		name    string
		args    args
		wantR   RemixBillingCSV
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				row: []string{
					"2022年11月分", "PP21XX", "0300XX", "ばしょ", "202", "5,123", "", "",
				},
			},
			wantR: RemixBillingCSV{
				BillingMonth:       "2022年11月分",
				ContractNumber:     "PP21XX",
				ProvidePointNumber: "0300XX",
				FacilityName:       "ばしょ",
				TotalConsumption:   202,
				Price:              "5,123",
			},
			wantErr: false,
		},
		{
			name: "invalid fields",
			args: args{
				row: []string{
					"2022年11月分", "PP21XX", "0300XX", "ばしょ", "202.", "5,123", "", "",
				},
			},
			wantR:   RemixBillingCSV{},
			wantErr: true,
		},
		{
			name: "invalid short field",
			args: args{
				row: []string{
					"2022年11月分", "PP21XX", "0300XX", "ばしょ", "202", "5,123", "",
				},
			},
			wantR:   RemixBillingCSV{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := NewRemixBillingCSV(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRemixBillingCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("NewRemixBillingCSV() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestRemixBillingCSV_ConvDBModel(t *testing.T) {
	type fields struct {
		BillingMonth       string
		ContractNumber     string
		ProvidePointNumber string
		FacilityName       string
		TotalConsumption   int
		Price              string
	}
	tests := []struct {
		name       string
		fields     fields
		wantRecord BillElect
		wantErr    bool
	}{
		{
			name: "ok",
			fields: fields{
				BillingMonth:       "2022年11月分",
				ContractNumber:     "PP21XX",
				ProvidePointNumber: "0300XX",
				FacilityName:       "ばしょ",
				TotalConsumption:   202,
				Price:              "5,123",
			},
			wantRecord: BillElect{
				BillingMonth:     "202211",
				TotalConsumption: 202000,
				Price:            5123,
			},
			wantErr: false,
		},
		{
			name: "invalid price",
			fields: fields{
				BillingMonth:       "2022年11月分",
				ContractNumber:     "PP21XX",
				ProvidePointNumber: "0300XX",
				FacilityName:       "ばしょ",
				TotalConsumption:   202,
				Price:              "5,x123",
			},
			wantRecord: BillElect{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RemixBillingCSV{
				BillingMonth:       tt.fields.BillingMonth,
				ContractNumber:     tt.fields.ContractNumber,
				ProvidePointNumber: tt.fields.ProvidePointNumber,
				FacilityName:       tt.fields.FacilityName,
				TotalConsumption:   tt.fields.TotalConsumption,
				Price:              tt.fields.Price,
			}
			gotRecord, err := r.ConvDBModel()
			if (err != nil) != tt.wantErr {
				t.Errorf("RemixBillingCSV.ConvDBModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("RemixBillingCSV.ConvDBModel() = %v, want %v", gotRecord, tt.wantRecord)
			}
		})
	}
}
