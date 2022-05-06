package db_ope

import "errors"

var (
	ErrRecordAlreadyExists = errors.New("this record is already recorded")
	ErrInternal            = errors.New("internal error")
	ErrConditionMismatch   = errors.New("database condition mismatch")
)
