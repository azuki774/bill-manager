package twitter_client

import (
	pb "github.com/azuki774/bill-manager/internal/grpcapi"
	"google.golang.org/grpc"
)

func MakeGrpcClient(conn *grpc.ClientConn) (client *pb.ElectConsumeServiceClient, err error) {
	c := pb.NewElectConsumeServiceClient(conn)
	return &c, nil
}
