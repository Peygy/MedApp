package grpc

import (
	"context"

	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	pb "github.com/peygy/medapp/internal/pkg/protos/graph_auth"
	"github.com/peygy/medapp/internal/services/auth_service/internal/managers"
)

func InitAuthGrpcServer(
	server *grpc.GrpcServer,
	roleManager managers.IRoleManager,
	userManager managers.IUserManager,
	logger logger.ILogger) {
	grpcServer := &grpcServer{
		roleManager: roleManager,
		userManager: userManager,
		log:         logger,
	}
	pb.RegisterSignUpServiceServer(server.Engine, grpcServer)

	logger.Info("Initialize of grpc server successfully")
}

type grpcServer struct {
	pb.UnimplementedSignUpServiceServer

	roleManager managers.IRoleManager
	userManager managers.IUserManager

	log logger.ILogger
}

func (s *grpcServer) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponce, error) {
	userRole := "user"
	user := managers.UserRecord{UserName: in.Username, Password: in.Password}
	userId, err := s.userManager.InsertUser(user)
	if err != nil {
		return nil, err
	}

	err = s.roleManager.AddRolesToUser(userId, userRole)
	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponce{UserId: userId, Role: userRole}, nil
}
