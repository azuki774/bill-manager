package repository

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"azuki774/bill-manager/internal/model"
)

type FileLoader struct{}

func (f *FileLoader) LoadRecordsFromJSON(ctx context.Context, filePath string) (recs []model.CreateRecord, err error) {
	content, err := os.Open(filePath)
	if err != nil {
		return recs, err
	}

	contentBin, err := ioutil.ReadAll(content)
	if err != nil {
		return recs, err
	}

	var incs []model.InCreateRecord
	err = json.Unmarshal(contentBin, &incs)
	if err != nil {
		return recs, err
	}

	for _, inc := range incs {
		var rec model.CreateRecord
		rec.FromInCreateRecord(ctx, &inc)
		recs = append(recs, rec)
	}
	return recs, nil
}
