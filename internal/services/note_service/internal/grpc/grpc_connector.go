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
	err := s.noteService.InsertUserNote(in.UserId, in.DoctorName, in.Specialization, in.VisitDate)
	if err != nil {
		return nil, err
	}

	return &pb.AddVisitRecordResponse{
		Success: true,
	}, nil
}

func (s *grpcServer) GetUserVisitRecords(ctx context.Context, in *pb.GetUserVisitRecordsRequest) (*pb.GetUserVisitRecordsResponse, error) {
	noteData, err := s.noteService.GetUserNotes(in.UserId)
	if err != nil {
		return nil, err
	}

	var visitRecords []*pb.VisitRecord
	for _, note := range noteData {
		visitRecord := &pb.VisitRecord{
			RecordNumber:   note.Id,
			DoctorName:     note.Doctor_name,
			Specialization: note.Specialization,
			VisitDate:      note.Visit_date,
		}
		visitRecords = append(visitRecords, visitRecord)
	}

	return &pb.GetUserVisitRecordsResponse{
		VisitRecords: visitRecords,
	}, nil
}
