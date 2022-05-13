package db_ope

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func makeDsn(dbAddr string, dbUserName string, dbUserPass string, dbName string) (dsn string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUserName, dbUserPass, dbAddr, dbName)
	return dsn
}

func DBConnect(dbAddr string, dbUserName string, dbUserPass string, dbName string) (*gorm.DB, error) {
	logger.Debugw("DB settings", "dbAddr", dbAddr, "dbUserName", dbUserName, "dbPass", dbUserPass, "dbName", dbName)
	dsn := makeDsn(dbAddr, dbUserName, dbUserPass, dbName)
	logger.Infow("make dsn", "dsn", dsn)
	ndb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorw("failed to connect database", "error", err)
		return nil, err
	}

	logger.Infow("connect to the database")
	return ndb, nil
}

func DBClose(db *gorm.DB) error {
	cdb, err := db.DB()
	if err != nil {
		logger.Errorw("database close error", "error", err)
		return err
	}

	err = cdb.Close()
	if err != nil {
		logger.Errorw("database close error", "error", err)
		return err
	}

	logger.Infow("database close")
	return nil
}
