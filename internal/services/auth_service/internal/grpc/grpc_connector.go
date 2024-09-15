package grpc

import (
	"context"

	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	pb "github.com/peygy/medapp/internal/pkg/protos/graph_auth"
	"github.com/peygy/medapp/internal/pkg/rabbitmq"
	"github.com/peygy/medapp/internal/services/auth_service/internal/managers"
)

func InitAuthGrpcServer(
	server *grpc.GrpcServer,
	passwordManager managers.IPasswordManager,
	roleManager managers.IRoleManager,
	userManager managers.IUserManager,
	rabbitMQServer rabbitmq.IRabbitMQServer,
	logger logger.ILogger,
) {
	grpcServer := &grpcServer{
		passwordManager: passwordManager,
		roleManager:     roleManager,
		userManager:     userManager,
		rabbitMQServer:  rabbitMQServer,
		log:             logger,
	}
	pb.RegisterAuthServiceServer(server.Engine, grpcServer)

	logger.Info("Initialize of grpc server successfully")
}

type grpcServer struct {
	pb.UnimplementedAuthServiceServer

	passwordManager managers.IPasswordManager
	roleManager     managers.IRoleManager
	userManager     managers.IUserManager
	rabbitMQServer  rabbitmq.IRabbitMQServer

	log logger.ILogger
}

func (s *grpcServer) SignUp(ctx context.Context, in *pb.UserRequest) (*pb.AuthResponce, error) {
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

	// publish message
	message := struct {
		UserId string `json:"userId"`
	}{UserId: userId}

	if err := s.rabbitMQServer.(message); err != nil {
		return nil, err
	}

	return &pb.AuthResponce{UserId: userId, Role: userRole}, nil
}

func (s *grpcServer) SignIn(ctx context.Context, in *pb.UserRequest) (*pb.AuthResponce, error) {
	user, err := s.userManager.GetUserByName(in.Username)
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

	return &pb.AuthResponce{UserId: user.Id, Role: userRole}, nil
}

func (s *grpcServer) GetUserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponce, error) {
	user, err := s.userManager.GetUserById(in.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.UserInfoResponce{UserName: user.UserName}, nil
}
