package api

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func LoadConf(l *zap.SugaredLogger) {
	logger = l
}
