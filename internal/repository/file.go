package repository

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"os"

	"azuki774/bill-manager/internal/model"
)

type FileLoader struct{}

func (f *FileLoader) LoadRecordsFromJSON(ctx context.Context, filePath string) (recs []model.CreateRecord, err error) {
	content, err := os.Open(filePath)
	if err != nil {
		return recs, err
	}

	contentBin, err := io.ReadAll(content)
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

func (f *FileLoader) LoadRemixElectConsumptionCSV(ctx context.Context, filePath string) (recs []model.RemixCSV, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []model.RemixCSV{}, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll() // csvを一度に全て読み込む
	if err != nil {
		return []model.RemixCSV{}, err
	}

	for i, row := range rows {
		// []string -> struct
		if i == 0 {
			// 1行目は 取得年月日,使用量合計(kWh),昼時間使用量(kWh),夜時間使用量(kWh) なので読み飛ばす
			continue
		}
		rec, err := model.NewRemixCSV(row)
		if err != nil && errors.Is(err, model.ErrNotProvided) {
			continue
		} else if err != nil {
			// internal error
			return []model.RemixCSV{}, err
		}

		recs = append(recs, rec)
	}

	return recs, nil
}

func (f *FileLoader) LoadRemixElectBillCSV(ctx context.Context, filePath string) (recs []model.RemixBillingCSV, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []model.RemixBillingCSV{}, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll() // csvを一度に全て読み込む
	if err != nil {
		return []model.RemixBillingCSV{}, err
	}

	for i, row := range rows {
		// []string -> struct
		if i == 0 {
			// 1行目は 請求年月,契約番号,供給地点特定番号,施設名称,使用量(kWh),請求金額(円),請求書PDF,電力使用CSV なので読み飛ばす
			continue
		}
		rec, err := model.NewRemixBillingCSV(row)
		if err != nil && errors.Is(err, model.ErrNotProvided) {
			continue
		} else if err != nil {
			// internal error
			return []model.RemixBillingCSV{}, err
		}

		recs = append(recs, rec)
	}

	return recs, nil
}

func (f *FileLoader) LoadWaterBillCSV(ctx context.Context, dir string) (recs []model.WaterBillingCSV, err error) {
	// TODO
	return recs, nil
}
