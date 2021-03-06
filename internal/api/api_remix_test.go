package api

import (
	"os"
	"reflect"
	"testing"
	"time"

	db "github.com/azuki774/bill-manager/internal/db-ope"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ElectConsumeDBRepositoryNormalGet struct {
	// すべて正常系
	db.UnimplementedElectConsumeDBRepository
}

func (dbR *ElectConsumeDBRepositoryNormalGet) GetElectConsume(tx *gorm.DB, t time.Time) (record db.ElectConsume, err error) {
	return db.ElectConsume{Id: 1, RecordDate: time.Date(2000, 1, 23, 0, 0, 0, 0, time.Now().Location()), Daytime: 100, Nighttime: 200, Total: 300}, nil
}

func (dbR *ElectConsumeDBRepositoryNormalGet) PostElectConsume(tx *gorm.DB, record db.ElectConsume) (err error) {
	return nil
}

func (dbR *ElectConsumeDBRepositoryNormalGet) GetCountElectConsume(tx *gorm.DB, t time.Time) (count int64, err error) {
	return 0, nil
}

func (dbR *ElectConsumeDBRepositoryNormalGet) mustEmbedUnimplementedElectConsumeDBRepository() {}

type ElectConsumeDBRepositoryErrorResponse struct {
	// すべてエラー応答
	db.UnimplementedElectConsumeDBRepository
}

func (dbR *ElectConsumeDBRepositoryErrorResponse) GetElectConsume(tx *gorm.DB, t time.Time) (record db.ElectConsume, err error) {
	return db.ElectConsume{}, gorm.ErrRecordNotFound
}

func (dbR *ElectConsumeDBRepositoryErrorResponse) PostElectConsume(tx *gorm.DB, record db.ElectConsume) (err error) {
	return gorm.ErrInvalidData
}

func (dbR *ElectConsumeDBRepositoryErrorResponse) GetCountElectConsume(tx *gorm.DB, t time.Time) (count int64, err error) {
	return 0, nil
}

func (dbR *ElectConsumeDBRepositoryErrorResponse) mustEmbedUnimplementedElectConsumeDBRepository() {}

type ElectConsumeDBRepositoryErrorExisted struct {
	// ErrDataAlreadyExisted
	db.UnimplementedElectConsumeDBRepository
}

func (dbR *ElectConsumeDBRepositoryErrorExisted) GetElectConsume(tx *gorm.DB, t time.Time) (record db.ElectConsume, err error) {
	return db.ElectConsume{Id: 1, RecordDate: time.Date(2000, 1, 23, 0, 0, 0, 0, time.Now().Location()), Daytime: 100, Nighttime: 200, Total: 300}, nil
}

func (dbR *ElectConsumeDBRepositoryErrorExisted) PostElectConsume(tx *gorm.DB, record db.ElectConsume) (err error) {
	return gorm.ErrInvalidData
}

func (dbR *ElectConsumeDBRepositoryErrorExisted) GetCountElectConsume(tx *gorm.DB, t time.Time) (count int64, err error) {
	return 1, nil
}

func (dbR *ElectConsumeDBRepositoryErrorExisted) mustEmbedUnimplementedElectConsumeDBRepository() {}

func setup() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	lg, _ := config.Build()
	defer lg.Sync() // flushes buffer, if any
	logger = lg.Sugar()
	LoadConf(logger)
}
func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	os.Exit(ret)
}

func Test_remixapiService_GetElectConsume(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name       string
		apis       *RemixapiServiceRepo
		args       args
		wantRecord db.ElectConsume
		wantErr    bool
	}{
		{
			name:       "Normally",
			apis:       &RemixapiServiceRepo{remixdbR: &ElectConsumeDBRepositoryNormalGet{}},
			args:       args{time.Date(2000, 1, 23, 0, 0, 0, 0, time.Now().Location())},
			wantRecord: db.ElectConsume{Id: 1, RecordDate: time.Date(2000, 1, 23, 0, 0, 0, 0, time.Now().Location()), Daytime: 100, Nighttime: 200, Total: 300},
			wantErr:    false,
		},
		{
			name:       "ErrNotFound",
			apis:       &RemixapiServiceRepo{remixdbR: &ElectConsumeDBRepositoryErrorResponse{}},
			args:       args{time.Date(2000, 1, 23, 0, 0, 0, 0, time.Now().Location())},
			wantRecord: db.ElectConsume{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecord, err := tt.apis.GetElectConsume(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("remixapiService.GetElectConsume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("remixapiService.GetElectConsume() = %v, want %v", gotRecord, tt.wantRecord)
			}
		})
	}
}

func TestRemixapiServiceRepo_PostElectConsume(t *testing.T) {
	type args struct {
		record db.ElectConsume
	}
	tests := []struct {
		name    string
		apis    *RemixapiServiceRepo
		args    args
		wantErr bool
	}{
		{
			name:    "Normally",
			apis:    &RemixapiServiceRepo{remixdbR: &ElectConsumeDBRepositoryNormalGet{}},
			args:    args{record: db.ElectConsume{RecordDate: time.Date(2000, 1, 23, 4, 5, 123, 123, time.Now().Location()), Daytime: 1000, Nighttime: 2000, Total: 3000}},
			wantErr: false,
		},
		{
			name:    "AlreadyExist",
			apis:    &RemixapiServiceRepo{remixdbR: &ElectConsumeDBRepositoryErrorExisted{}},
			args:    args{record: db.ElectConsume{RecordDate: time.Date(2000, 1, 23, 4, 5, 123, 123, time.Now().Location()), Daytime: 1000, Nighttime: 2000, Total: 3000}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.apis.PostElectConsume(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("RemixapiServiceRepo.PostElectConsume() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
