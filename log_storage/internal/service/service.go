package service

import (
	"context"
	"fmt"
	log "fmt"
	"strings"

	"github.com/imyazip/GoSIEM/log-storage/internal/storage"
	"github.com/imyazip/GoSIEM/log-storage/pkg/config"
	pb "github.com/imyazip/GoSIEM/log-storage/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type LogService struct {
	config     config.Config
	storage    *storage.Storage
	authClient pb.AuthServiceClient
}

func NewLogService(db *storage.Storage, config config.Config) *LogService {
	if config.AuthServer.Port == 0 {
		panic(log.Sprintf("Invalid port number for Auth service: %d", config.AuthServer.Port))
	}
	authServAddr := log.Sprintf("%s:%d", config.AuthServer.Host, config.AuthServer.Port)
	log.Printf("Connecting to auth srv: %s", authServAddr)
	conn, err := grpc.NewClient(authServAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(log.Sprint("failed to connect to auth service:", authServAddr))
	}

	authClient := pb.NewAuthServiceClient(conn)

	return &LogService{
		config:     config,
		storage:    db,
		authClient: authClient,
	}
}

func (s *LogService) ValidateJWT(ctx context.Context) (bool, error) {
	// Получаем метаданные из контекста
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, fmt.Errorf("missing metadata in context")
	}

	// Получаем значение из заголовка "authorization"
	authHeader := md.Get("authorization")
	if len(authHeader) == 0 {
		return false, fmt.Errorf("missing JWT token")
	}

	// Проверяем, что строка начинается с "Bearer "
	if !strings.HasPrefix(authHeader[0], "Bearer ") {
		return false, fmt.Errorf("invalid token format")
	}

	// Извлекаем сам токен (удаляем "Bearer " префикс)
	token := strings.TrimPrefix(authHeader[0], "Bearer ")

	// Проверяем валидность JWT через authClient
	resp, err := s.authClient.ValidateJWTForSensor(ctx, &pb.ValidateJWTForSensorRequest{JWT: token})
	if err != nil {
		return false, fmt.Errorf("failed to validate JWT: %v", err)
	}

	if !resp.Valid {
		return false, fmt.Errorf("invalid JWT token")
	}

	return true, nil
}

// SaveRawLog сохраняет сырой строковый лог в базу
func (s *LogService) SaveRawLog(ctx context.Context, req *pb.TransferRawStringLogRequest) error {
	return s.storage.InsertRawLog(req.LogSource, req.LogString, req.SystemCreatedAt.AsTime(), req.SensorId)
}

// SaveSerializedLog сохраняет сериализованные логи в базу
func (s *LogService) SaveSerializedLog(ctx context.Context, req *pb.TranserSerializedLogRequest) error {
	SerializedLogsArray := &pb.StringArray{Items: req.LogSerialized} //Создаем вспомогательную структуру для сериализации массива из сообщения TranserSerializedLogRequest

	serializedData, err := proto.Marshal(SerializedLogsArray) //Конвертируем в бинарный формат
	if err != nil {
		return log.Errorf("failed to serialize log data: %w", err)
	}
	return s.storage.InsertSerializedLog(req.LogSource, serializedData, req.SystemCreatedAt.AsTime(), req.SensorId)
}
