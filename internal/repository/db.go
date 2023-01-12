package repository

import (
	"azuki774/bill-manager/internal/model"
	"errors"

	"gorm.io/gorm"
)

type DBRepository struct {
	Conn *gorm.DB
}

func (d *DBRepository) CloseDB() (err error) {
	dbconn, err := d.Conn.DB()
	if err != nil {
		return err
	}
	return dbconn.Close()
}

// AddElectConsumption inserts elect_consumption without overwriting.
func (d *DBRepository) AddElectConsumption(record model.ElectConsumption) (err error) {
	recordDate := record.RecordDate
	err = d.Conn.Where("record_date = ?", recordDate).Take(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// not found -> ok
	} else if err != nil {
		// internal error
		return err
	} else {
		// record exists
		return nil
	}

	err = d.Conn.Create(&record).Error
	if err != nil {
		return err
	}

	return nil
}

// AddElectBill inserts bill_elect without overwriting.
func (d *DBRepository) AddElectBill(record model.BillElect) (err error) {
	bm := record.BillingMonth
	err = d.Conn.Where("billing_month = ?", bm).Take(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// not found -> ok
	} else if err != nil {
		// internal error
		return err
	} else {
		// record exists
		return nil
	}

	err = d.Conn.Create(&record).Error
	if err != nil {
		return err
	}

	return nil
}

// AddWaterBill inserts bill_water without overwriting.
func (d *DBRepository) AddWaterBill(r model.BillWater) (err error) {
	bm := r.BillingMonth
	err = d.Conn.Where("billing_month = ?", bm).Take(&r).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// not found -> ok
	} else if err != nil {
		// internal error
		return err
	} else {
		// record exists
		return nil
	}

	err = d.Conn.Create(&r).Error
	if err != nil {
		return err
	}

	return nil
}

// AddBillGas inserts bill_gas without overwriting.
func (d *DBRepository) AddGasBill(r model.BillGas) (err error) {
	bm := r.BillingMonth
	err = d.Conn.Where("billing_month = ?", bm).Take(&r).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// not found -> ok
	} else if err != nil {
		// internal error
		return err
	} else {
		// record exists
		return nil
	}

	err = d.Conn.Create(&r).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *DBRepository) GetBillElect(yyyymm string) (b model.BillElect, err error) {
	err = d.Conn.Table(model.BillElect{}.TableName()).Where("billing_month = ?", yyyymm).Take(&b).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.BillElect{}, model.ErrNotFound
	}
	if err != nil {
		return model.BillElect{}, nil
	}
	return b, nil
}

func (d *DBRepository) GetBillGas(yyyymm string) (b model.BillGas, err error) {
	err = d.Conn.Table(model.BillGas{}.TableName()).Where("billing_month = ?", yyyymm).Take(&b).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.BillGas{}, model.ErrNotFound
	}
	if err != nil {
		return model.BillGas{}, nil
	}
	return b, nil
}

func (d *DBRepository) GetBillWater(yyyymm string) (b model.BillWater, err error) {
	err = d.Conn.Table(model.BillWater{}.TableName()).Where("billing_month = ?", yyyymm).Take(&b).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.BillWater{}, model.ErrNotFound
	}
	if err != nil {
		return model.BillWater{}, nil
	}
	return b, nil
}
