package grpc

import (
	"context"

	"github.com/peygy/medapp/internal/pkg/grpc"
	"github.com/peygy/medapp/internal/pkg/logger"
	pb "github.com/peygy/medapp/internal/pkg/protos/graph_crew"
	"github.com/peygy/medapp/internal/services/crew_service/internal/services"
)

func InitCrewGrpcServer(
	server *grpc.GrpcServer,
	logger logger.ILogger,
	crewService services.ICrewService) {
	grpcServer := &grpcServer{
		crewService: crewService,
		log:         logger,
	}
	pb.RegisterCrewServiceServer(server.Engine, grpcServer)

	logger.Info("Initialize of grpc server successfully")
}

type grpcServer struct {
	pb.UnimplementedCrewServiceServer

	crewService services.ICrewService
	log         logger.ILogger
}

func (s *grpcServer) GetDoctors(ctx context.Context, in *pb.GetDoctorsRequest) (*pb.GetDoctorsResponse, error) {
	crewData, err := s.crewService.GetAllCrew()
	if err != nil {
		return nil, err
	}

	var doctors []*pb.Doctor
	for _, crew := range crewData {
		doctor := &pb.Doctor{
			DoctorId:        crew.DoctorId,
			DoctorName:      crew.DoctorName,
			Specialization:  crew.Specialization,
			ExperienceYears: int32(crew.ExperienceYears),
		}
		doctors = append(doctors, doctor)
	}

	return &pb.GetDoctorsResponse{
		Doctors: doctors,
	}, nil
}
