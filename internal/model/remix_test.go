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
