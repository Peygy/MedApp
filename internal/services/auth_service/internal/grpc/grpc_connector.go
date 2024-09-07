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
	passwordManager managers.IPasswordManager,
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

	passwordManager managers.IPasswordManager
	roleManager     managers.IRoleManager
	userManager     managers.IUserManager

	log logger.ILogger
}

func (s *grpcServer) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponce, error) {
	if err := s.passwordManager.ValidPassword(in.Password); err != nil {
		return nil, err
	}

	hashedPassword, err := s.passwordManager.HashPassword(in.Password)
	if err != nil {
		return nil, err
	}

	userRole := "user"
	user := managers.UserRecord{UserName: in.Username, Password: hashedPassword}
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

func (s *grpcServer) SignIn(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponce, error) {
	user, err := s.userManager.GetUser(in.Username)
	if err != nil {
		return nil, err
	}

	if err := s.passwordManager.CheckPasswordHash(in.Password, user.Password); err != nil {
		return nil, err
	}

	userRole, err := s.roleManager.GetUserRole(user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponce{UserId: user.Id, Role: userRole}, nil
}
