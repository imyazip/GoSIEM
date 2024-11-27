package server

import (
	"context"

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

func (h *LogStorageApi) TransferSerializedStringLog(ctx context.Context, req *pb.TranserSerializedLogRequest) (*pb.TranserSerializedLogResponse, error) {
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
