package server

import (
	"context"

	pb "github.com/imyazip/GoSIEM/log-storage/proto"
)

type LogStorageApi struct {
	pb.UnimplementedLogStorageServiceServer
	service *pb.LogStorageServiceServer
}

func NewLogStorageApi(service *pb.LogStorageServiceServer) *LogStorageApi {
	return &LogStorageApi{service: service}
}

func (h *LogStorageApi) TransferRawStringLog(ctx context.Context, req *pb.TransferRawStringLogRequest) (*pb.TransferRawStringLogResponse, error) {
	return nil, nil
}

func (h *LogStorageApi) TransferSerializedStringLog(ctx context.Context, req *pb.TranserSerializedLogRequest) (*pb.TranserSerializedLogResponse, error) {
	return nil, nil
}
