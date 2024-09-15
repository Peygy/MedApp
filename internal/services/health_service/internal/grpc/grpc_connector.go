package grpc

import (
	"context"

	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	pb "github.com/peygy/medapp/internal/pkg/protos/graph_health"
	"github.com/peygy/medapp/internal/services/health_service/internal/services"
)

func InitHealhGrpcServer(
	server *grpc.GrpcServer,
	logger logger.ILogger,
	healthService services.IHealthService) {
	grpcServer := &grpcServer{
		healthService: healthService,
		log:           logger,
	}
	pb.RegisterHealthServiceServer(server.Engine, grpcServer)

	logger.Info("Initialize of grpc server successfully")
}

type grpcServer struct {
	pb.UnimplementedHealthServiceServer

	healthService services.IHealthService
	log           logger.ILogger
}

func calcDailyWater(weight float32) float32 {
	return weight * 3
}

func calcBodyMassIndex(weight float32, height float32) float32 {
	height /= 100
	return weight / (height * height)
}

func (s *grpcServer) GetHealthData(ctx context.Context, in *pb.GetHealthDataRequest) (*pb.HealthDataResponce, error) {
	if err := s.healthService.InsertHealthData(in.UserId); err != nil {
		return nil, err
	}

	healthData, err := s.healthService.GetHealthDataByUserId(in.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.HealthDataResponce{
		Age:           healthData.Age,
		Height:        healthData.Height,
		Weight:        healthData.Weight,
		Pulse:         healthData.Pulse,
		Pressure:      healthData.Pressure,
		DailyWater:    calcDailyWater(healthData.Weight),
		BodyMassIndex: calcBodyMassIndex(healthData.Weight, healthData.Height),
	}, nil
}

func (s *grpcServer) UpdateHealthData(ctx context.Context, in *pb.UpdateHealthDataRequest) (*pb.HealthDataResponce, error) {
	oldHealthData := services.HealthData{
		Age:      in.Age,
		Height:   in.Height,
		Weight:   in.Weight,
		Pulse:    in.Pulse,
		Pressure: in.Pressure,
	}

	healthData, err := s.healthService.UpdateHealthDataByUserId(in.UserId, oldHealthData)
	if err != nil {
		return nil, err
	}

	return &pb.HealthDataResponce{
		Age:           healthData.Age,
		Height:        healthData.Height,
		Weight:        healthData.Weight,
		Pulse:         healthData.Pulse,
		Pressure:      healthData.Pressure,
		DailyWater:    calcDailyWater(healthData.Weight),
		BodyMassIndex: calcBodyMassIndex(healthData.Weight, healthData.Height),
	}, nil
}
