package main

import (
	"github.com/azuki774/bill-manager/internal/api"
	db "github.com/azuki774/bill-manager/internal/db-ope"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func main() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel) //出力するログレベルをDebugレベルに変更
	lg, _ := config.Build()
	// lg, _ := zap.NewProduction()
	defer lg.Sync() // flushes buffer, if any
	logger = lg.Sugar()
	api.LoadConf(logger)
	db.LoadConf(logger)

	Execute()
}
