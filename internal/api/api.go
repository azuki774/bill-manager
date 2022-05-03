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
	logger.Info("receive data")
	return &empty.Empty{}, nil
}

// func (e *server) ElectConsumeGet(context.Context, *grpcapi.Date) (*grpcapi.OnedayElectConsume, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method ElectConsumePost not implemented")
// }
