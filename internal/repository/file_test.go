package repository

import (
	"reflect"
	"testing"

	"github.com/azuki774/bill-manager/internal/model"
)

func TestFileLoader_LoadRecordsFromJSON(t *testing.T) {
	type args struct {
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
				filePath: "../../test/data.json",
			},
			f: &FileLoader{},
			wantRecs: []model.CreateRecord{
				{
					CategoryID: 100,
					Price:      123,
				},
				{
					CategoryID: 110,
					Price:      456,
					Memo:       "memomemo",
				},
				{
					CategoryID: 200,
					Price:      789,
					Date:       "20060102",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileLoader{}
			gotRecs, err := f.LoadRecordsFromJSON(tt.args.filePath)
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
