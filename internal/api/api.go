package api

import (
	"context"

	pb "github.com/azuki774/bill-manager/internal/grpcapi"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type server struct {
	pb.UnimplementedElectConsumeServiceServer
}

func (s *server) ElectConsumePost(ctx context.Context, in *pb.OnedayElectConsume) (*empty.Empty, error) {
	logger.Infow("receive data", "api", "ElectConsumePost")
	return &empty.Empty{}, nil
}

func (e *server) ElectConsumeGet(context.Context, *pb.DateStruct) (*pb.OnedayElectConsume, error) {
	logger.Infow("receive data", "api", "ElectConsumeGet")
	retData := pb.OnedayElectConsume{Daytime: 1.0, Nighttime: 2.0, Total: 3.0}
	return &retData, nil
	// return nil, status.Errorf(codes.Unimplemented, "method ElectConsumePost not implemented")
}
