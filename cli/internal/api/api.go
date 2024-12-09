package api

import (
	"context"
	"fmt"

	pb "github.com/imyazip/GoSIEM/cli/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// Функция для получения токена
func Login(username, password string, client pb.AuthServiceClient) (string, error) {
	// Создаем запрос с API-ключом
	req := &pb.LoginRequest{
		Username: username,
		Password: password,
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
