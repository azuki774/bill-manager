package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/azuki774/bill-manager/internal/model"
)

type FileLoader struct{}

func (f *FileLoader) LoadRecordsFromJSON(filePath string) (recs []model.CreateRecord, err error) {
	content, err := os.Open(filePath)
	if err != nil {
		return recs, err
	}

	contentBin, err := ioutil.ReadAll(content)
	if err != nil {
		return recs, err
	}

	err = json.Unmarshal(contentBin, &recs)
	if err != nil {
		return recs, err
	}
	return recs, nil
}
