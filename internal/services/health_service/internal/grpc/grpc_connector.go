package grpc

import (
	"context"

	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	pb "github.com/peygy/medapp/internal/pkg/protos/graph_auth"
)

func InitAuthGrpcServer(
	server *grpc.GrpcServer,
	logger logger.ILogger) {
	grpcServer := &grpcServer{
		log:             logger,
	}
	pb.RegisterAuthServiceServer(server.Engine, grpcServer)

	logger.Info("Initialize of grpc server successfully")
}

type grpcServer struct {
	pb.UnimplementedAuthServiceServer

	log logger.ILogger
}

func (s *grpcServer) GetHealthData(ctx context.Context, in *pb.UserRequest) (*pb.AuthResponce, error) {

	return &pb.AuthResponce{UserId: userId, Role: userRole}, nil
}

func (s *grpcServer) UpdateHealthData(ctx context.Context, in *pb.UserRequest) (*pb.AuthResponce, error) {

	return &pb.AuthResponce{UserId: user.Id, Role: userRole}, nil
}