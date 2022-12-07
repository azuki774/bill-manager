package repository

import (
	"azuki774/bill-manager/internal/model"
	"context"
	"reflect"
	"testing"
	"time"
)

func TestFileLoader_LoadRecordsFromJSON(t *testing.T) {
	type args struct {
		ctx      context.Context
		filePath string
	}
	tests := []struct {
		name     string
		f        *FileLoader
		args     args
		wantRecs []model.CreateRecord
		wantErr  bool
	}{
		{
			name: "ok",
			args: args{
				ctx:      model.NewCtxYYYYMM(time.Date(2000, 5, 1, 0, 0, 0, 0, time.Local)),
				filePath: "../../test/data.json",
			},
			f: &FileLoader{},
			wantRecs: []model.CreateRecord{
				{
					CategoryID: 100,
					Price:      123,
					From:       "bill-manager-mawinter",
				},
				{
					CategoryID: 110,
					Price:      456,
					From:       "bill-manager-mawinter",
					Type:       "S",
					Memo:       "memomemo",
				},
				{
					CategoryID: 200,
					Price:      789,
					Date:       "20000505",
					From:       "bill-manager-mawinter",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileLoader{}
			gotRecs, err := f.LoadRecordsFromJSON(tt.args.ctx, tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileLoader.LoadRecordsFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecs, tt.wantRecs) {
				t.Errorf("FileLoader.LoadRecordsFromJSON() = %v, want %v", gotRecs, tt.wantRecs)
			}
		})
	}
}

func TestFileLoader_LoadRemixElectConsumptionCSV(t *testing.T) {
	type args struct {
		ctx      context.Context
		filePath string
	}
	tests := []struct {
		name     string
		f        *FileLoader
		args     args
		wantRecs []model.RemixCSV
		wantErr  bool
	}{
		{
			name: "ok",
			args: args{
				ctx:      context.Background(),
				filePath: "../../test/data.csv",
			},
			f: &FileLoader{},
			wantRecs: []model.RemixCSV{
				{
					RecordDate:       "2022/12/02",
					TotalConsumption: 6,
					DayConsumption:   2,
					NightConsumption: 4,
				},
				{
					RecordDate:       "2022/12/01",
					TotalConsumption: 6,
					DayConsumption:   4,
					NightConsumption: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileLoader{}
			gotRecs, err := f.LoadRemixElectConsumptionCSV(tt.args.ctx, tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileLoader.LoadRemixElectConsumptionCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecs, tt.wantRecs) {
				t.Errorf("FileLoader.LoadRemixElectConsumptionCSV() = %v, want %v", gotRecs, tt.wantRecs)
			}
		})
	}
}
