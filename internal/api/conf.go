package api

import (
	db "github.com/azuki774/bill-manager/internal/db-ope"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger
var dbR db.ElectConsumeDBRepository
var apis RemixapiServiceRepository

func LoadConf(l *zap.SugaredLogger) {
	logger = l
}
