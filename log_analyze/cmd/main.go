package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/imyazip/GoSIEM/log_analzye/config"
	pb "github.com/imyazip/GoSIEM/log_analzye/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

func deserializeSerializedLog(data []byte) (*pb.StringArray, error) {
	// Создаем экземпляр структуры, куда будут десериализованы данные
	deserializedData := &pb.StringArray{}

	// Выполняем десериализацию
	err := proto.Unmarshal(data, deserializedData)
	if err != nil {
		return nil, err
	}

	return deserializedData, nil
}

func main() {
	cfg := config.LoadConfig("analyzer.yaml")
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.LogServer.Host, cfg.LogServer.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewLogStorageServiceClient(conn)

	for {
		req := &pb.GetNewLogsRequest{
			Limit: 10,
		}
		logResp, err := pb.LogStorageServiceClient.GetNewLogs(client, context.Background(), req)
		if err != nil {
			log.Printf("Failed to get new logs: %s", err)
		}
		logs := logResp.GetLogs()
		for _, logEntry := range logs {
			deserialized, err := deserializeSerializedLog(logEntry.LogSerialized)
			if err != nil {
				log.Printf("Error deserializing log: %s", err)
			}

			for _, logContains := range deserialized.Items {
				log.Print(logContains)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
