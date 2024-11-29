package api

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/imyazip/GoSIEM/windows_agent/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Функция для получения токена
func GetToken(apiKey string, client pb.AuthServiceClient) (string, error) {
	// Создаем запрос с API-ключом
	req := &pb.GenerateJWTForSensorRequest{
		ApiKey: apiKey,
	}

	// Отправляем запрос и получаем ответ
	res, err := client.GenerateJWTForSensor(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("error during request: %v", err)
	}

	return res.Token, nil
}

func GetAuthClient(conn *grpc.ClientConn) pb.AuthServiceClient {
	return pb.NewAuthServiceClient(conn)
}

func GetLogClient(conn *grpc.ClientConn) pb.LogStorageServiceClient {
	return pb.NewLogStorageServiceClient(conn)
}

func ConnectToServer(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials())) // Замените на актуальный адрес
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	return conn, nil
}

func CreateAuthContext(token string) context.Context {
	return metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", "Bearer "+token))
}

func SendSerializedLog(ctx context.Context, serialized []string, source string, sensor_id string, logClient pb.LogStorageServiceClient) error {
	req := &pb.TranserSerializedLogRequest{
		LogSource:       "example_source",
		LogSerialized:   serialized,
		SystemCreatedAt: timestamppb.New(time.Now()),
		SensorId:        "testsensor", // Пример SensorId
	}

	// Отправляем запрос
	resp, err := logClient.TranserSerializedLog(ctx, req)
	if err != nil {
		return nil
	}
	// Проверяем ответ
	if resp.Success {
		log.Printf("Successfully sent log")
	} else {
		log.Printf("Server failed to process log: %s", resp.Error)
	}
	return nil
}
