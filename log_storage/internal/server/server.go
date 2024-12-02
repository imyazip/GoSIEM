package server

import (
	"context"
	"log"

	"github.com/imyazip/GoSIEM/log-storage/internal/service"
	pb "github.com/imyazip/GoSIEM/log-storage/proto"
)

type LogStorageApi struct {
	pb.UnimplementedLogStorageServiceServer
	service *service.LogService
}

func NewLogStorageApi(service *service.LogService) *LogStorageApi {
	return &LogStorageApi{service: service}
}

func (h *LogStorageApi) TransferRawStringLog(ctx context.Context, req *pb.TransferRawStringLogRequest) (*pb.TransferRawStringLogResponse, error) {
	_, err := h.service.ValidateJWT(ctx)
	if err != nil {
		return &pb.TransferRawStringLogResponse{Success: false, Error: err.Error()}, nil
	}
	err = h.service.SaveRawLog(ctx, req)
	if err != nil {
		return &pb.TransferRawStringLogResponse{Success: false, Error: err.Error()}, nil
	}

	return &pb.TransferRawStringLogResponse{Success: true}, nil
}

func (h *LogStorageApi) TranserSerializedLog(ctx context.Context, req *pb.TranserSerializedLogRequest) (*pb.TranserSerializedLogResponse, error) {
	log.Println("Received request to TranserSerializedLog")
	_, err := h.service.ValidateJWT(ctx)
	if err != nil {
		return &pb.TranserSerializedLogResponse{Success: false, Error: err.Error()}, nil
	}
	err = h.service.SaveSerializedLog(ctx, req)
	if err != nil {
		return &pb.TranserSerializedLogResponse{Success: false, Error: err.Error()}, nil
	}

	return &pb.TranserSerializedLogResponse{Success: true}, nil
}

func (h *LogStorageApi) GetNewLogs(ctx context.Context, req *pb.GetNewLogsRequest) (*pb.GetNewLogsResponse, error) {
	answer, err := h.service.GetNewLogs(ctx, req)
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (h *LogStorageApi) AddSecurityEvent(ctx context.Context, req *pb.AddSecurityEventRequest) (*pb.AddSecurityEventResponse, error) {
	answer, err := h.service.AddSecurityEvent(ctx, req)
	if err != nil {
		return nil, err
	}

	return answer, nil
}
