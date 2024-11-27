package service

import (
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
