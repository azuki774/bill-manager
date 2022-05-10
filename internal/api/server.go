package api

import (
	"context"
	"net"
	"time"

	db "github.com/azuki774/bill-manager/internal/db-ope"
	pb "github.com/azuki774/bill-manager/internal/grpcapi"
	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedElectConsumeServiceServer
}

func Start() (err error) {
	// Set Database Connector
	time.Sleep(20 * time.Second) // for DB preparing
	dbconn, err := db.DBConnect("bill-manager-db", "root", "billmanager", "BILLMANAGER")
	if err != nil {
		return err
	}
	defer db.DBClose(dbconn)

	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		logger.Fatalw("grpc server establish error", "error", err)
		return err
	}

	s := grpc.NewServer()
	svc := server{}
	pb.RegisterElectConsumeServiceServer(s, &svc)

	logger.Info("grpc server start")
	reflection.Register(s)

	dbR = db.NewElectConsumeDBRepository(dbconn)
	apis = NewRemixapiService(dbR)

	if err := s.Serve(lis); err != nil {
		logger.Infow("grpc server end", "error", err)
		return err
	}

	logger.Infow("grpc server end")
	return nil
}

func (s *server) ElectConsumePost(ctx context.Context, in *pb.OnedayElectConsume) (*empty.Empty, error) {
	logger.Infow("receive data", "api", "ElectConsumePost", "data", in)
	return &empty.Empty{}, nil
}

func (s *server) ElectConsumeGet(ctx context.Context, in *pb.DateStruct) (*pb.OnedayElectConsume, error) {
	logger.Infow("receive data", "api", "ElectConsumeGet")

	t := time.Date(int(in.Year), time.Month(int(in.Month)), int(in.Day), 0, 0, 0, 0, time.Now().Location())
	record, err := apis.GetElectConsume(t)
	if err != nil {
		logger.Error("error", err)
	}
	retData := pb.OnedayElectConsume{Daytime: record.Daytime, Nighttime: record.Nighttime, Total: record.Total}
	return &retData, nil
}
