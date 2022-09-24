package main

import (
	"github.com/azuki774/bill-manager/internal/api"
	db "github.com/azuki774/bill-manager/internal/repository"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func main() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel) // Change to DebugLevel
	lg, _ := config.Build()
	defer lg.Sync() // flushes buffer, if any
	logger = lg.Sugar()
	api.LoadConf(logger)
	db.LoadConf(logger)

	Execute()
}
