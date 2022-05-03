package main

import (
	twitter_client "github.com/azuki774/bill-manager/internal/twitter-client"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func main() {
	lg, _ := zap.NewProduction()
	defer lg.Sync() // flushes buffer, if any
	logger = lg.Sugar()
	twitter_client.LoadConf(logger)

	Execute()
}
