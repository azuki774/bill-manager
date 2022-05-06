package main

import (
	"github.com/azuki774/bill-manager/internal/api"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func main() {
	lg, _ := zap.NewProduction()
	defer lg.Sync() // flushes buffer, if any
	logger = lg.Sugar()
	api.LoadConf(logger)

	Execute()
}
