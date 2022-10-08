package model

import (
	"context"
	"time"
)

type yyyymm struct{}

var yyyymmKey yyyymm

// GetYYYYMM gets yyyymmKey from ctx.
// If not set, return yyyymm.
func GetYYYYMM(ctx context.Context) string {
	v := ctx.Value(yyyymmKey)

	yyyymm, ok := v.(string)
	if !ok {
		return time.Now().Format("200601")
	}

	return yyyymm
}

// For test
func NewCtxYYYYMM(t time.Time) context.Context {
	ctx := context.WithValue(context.Background(), yyyymmKey, t.Format("200601"))
	return ctx
}
