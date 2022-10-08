package model

import (
	"context"
	"testing"
	"time"
)

func TestGetYYYYMM(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "exist",
			args: args{
				ctx: NewCtxYYYYMM(time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)),
			},
			want: "200001",
		},
		{
			name: "not exist",
			args: args{
				ctx: context.Background(),
			},
			want: time.Now().Format("200601"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetYYYYMM(tt.args.ctx); got != tt.want {
				t.Errorf("GetYYYYMM() = %v, want %v", got, tt.want)
			}
		})
	}
}
