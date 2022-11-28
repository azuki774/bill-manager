package repository

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/azuki774/bill-manager/internal/model"
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
