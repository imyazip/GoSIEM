package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/imyazip/GoSIEM/log-storage/internal/storage"
	"github.com/imyazip/GoSIEM/log-storage/pkg/config"
	pb "github.com/imyazip/GoSIEM/log-storage/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LogService struct {
	config     config.Config
	storage    *storage.Storage
	authClient pb.AuthServiceClient
}

func NewLogService(db *storage.Storage, config config.Config) *LogService {
	authServAddr := fmt.Sprintf("%s:%d", config.AuthServer.Host, config.AuthServer.Port)
	conn, err := grpc.NewClient(authServAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprint("failed to connect to auth service:", authServAddr))
	}

	authClient := pb.NewAuthServiceClient(conn)

	return &LogService{
		config:     config,
		storage:    db,
		authClient: authClient,
	}
}

func (s *LogService) ValidateJWT(ctx context.Context) (bool, error) {
	jwt, ok := ctx.Value("jwt").(string)
	if !ok || jwt == "" {
		return false, fmt.Errorf("missing JWT token")
	}

	resp, err := s.authClient.ValidateJWTKey(ctx, &pb.ValidateJWTKeyRequest{JWT: jwt})
	if err != nil {
		return false, fmt.Errorf("failed to validate JWT: %w", err)
	}

	if !resp.Valid {
		return false, fmt.Errorf("invalid JWT token")
	}

	return true, nil
}

// SaveRawLog сохраняет сырой строковый лог в хранилище
func (s *LogService) SaveRawLog(ctx context.Context, req *pb.TransferRawStringLogRequest, sensorID int64) error {
	return s.storage.InsertRawLog(req.LogSource, req.LogString, req.SystemCreatedAt.AsTime(), req.SensorId)
}

// SaveSerializedLog сохраняет сериализованные логи в хранилище
func (s *LogService) SaveSerializedLog(ctx context.Context, req *pb.TranserSerializedLogRequest, sensorID int64) error {
	serializedData, err := json.Marshal(req.LogSerialized)
	if err != nil {
		return fmt.Errorf("failed to serialize log data: %w", err)
	}
	return s.storage.InsertSerializedLog(req.LogSource, string(serializedData), req.SystemCreatedAt.AsTime(), fmt.Sprintf("%d", sensorID))
}
