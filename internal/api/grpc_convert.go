package api

import (
	"time"

	pb "github.com/azuki774/bill-manager/internal/grpcapi"
)

func grpcDateStructTodateTime(date *pb.DateStruct) (t time.Time) {
	t = time.Date(int(date.Year), time.Month(int(date.Month)), int(date.Day), 0, 0, 0, 0, time.Now().Location())
	return t
}
