package grpc

import (
	"context"

	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	pb "github.com/peygy/medapp/internal/pkg/protos/graph_note"
	"github.com/peygy/medapp/internal/services/note_service/internal/services"
)

func InitNoteGrpcServer(
	server *grpc.GrpcServer,
	logger logger.ILogger,
	noteService services.INoteService) {
	grpcServer := &grpcServer{
		noteService: noteService,
		log:         logger,
	}
	pb.RegisterVisitServiceServer(server.Engine, grpcServer)

	logger.Info("Initialize of grpc server successfully")
}

type grpcServer struct {
	pb.UnimplementedVisitServiceServer

	noteService services.INoteService
	log         logger.ILogger
}

func (s *grpcServer) AddVisitRecord(ctx context.Context, in *pb.AddVisitRecordRequest) (*pb.AddVisitRecordResponse, error) {
	healthData, err := s.noteService.InsertUserNote(in.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.AddVisitRecordResponse{
		Age:           healthData.Age,
		Height:        healthData.Height,
		Weight:        healthData.Weight,
		Pulse:         healthData.Pulse,
		Pressure:      healthData.Pressure,
		DailyWater:    calcDailyWater(healthData.Weight),
		BodyMassIndex: calcBodyMassIndex(healthData.Weight, healthData.Height),
	}, nil
}

func (s *grpcServer) GetUserVisitRecords(ctx context.Context, in *pb.GetUserVisitRecordsRequest) (*pb.GetUserVisitRecordsResponse, error) {
	oldHealthData := services.HealthData{
		Age:      in.Age,
		Height:   in.Height,
		Weight:   in.Weight,
		Pulse:    in.Pulse,
		Pressure: in.Pressure,
	}

	healthData, err := s.noteService.GetUserNotes(in.UserId, oldHealthData)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserVisitRecordsResponse{
		Age:           healthData.Age,
		Height:        healthData.Height,
		Weight:        healthData.Weight,
		Pulse:         healthData.Pulse,
		Pressure:      healthData.Pressure,
		DailyWater:    calcDailyWater(healthData.Weight),
		BodyMassIndex: calcBodyMassIndex(healthData.Weight, healthData.Height),
	}, nil
}
