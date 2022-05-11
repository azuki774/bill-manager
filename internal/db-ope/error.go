package db_ope

import "errors"

var (
	ErrInternal            = errors.New("internal error")
	ErrConditionMismatch   = errors.New("database condition mismatch")
	ErrRecordAlreadyExists = errors.New("record is already existed")
)
