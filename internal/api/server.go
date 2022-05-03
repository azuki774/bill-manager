package api

import (
	"net"

	pb "github.com/azuki774/bill-manager/internal/grpcapi"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() (err error) {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		logger.Fatal("grpc server establish error", zap.Error(err))
		return err
	}

	s := grpc.NewServer()
	svc := server{}
	pb.RegisterElectConsumeServiceServer(s, &svc)

	logger.Info("grpc server start")
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logger.Info("grpc server end", zap.Error(err))
		return err
	}

	logger.Info("grpc server end")
	return nil
}
