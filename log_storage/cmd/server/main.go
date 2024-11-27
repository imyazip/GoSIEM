package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	handler "github.com/imyazip/GoSIEM/log-storage/internal/server"
	"github.com/imyazip/GoSIEM/log-storage/internal/service"
	"github.com/imyazip/GoSIEM/log-storage/internal/storage"
	"github.com/imyazip/GoSIEM/log-storage/middleware"
	"github.com/imyazip/GoSIEM/log-storage/pkg/config"
	pb "github.com/imyazip/GoSIEM/log-storage/proto"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig("server-config.yaml")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name,
	)

	// Подключаемся к базе данных
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err) // Логируем ошибку, если не удалось подключиться
	}
	defer db.Close()

	storage := storage.NewStorage(db)

	// Контекст с таймаутом для попыток выполнения миграции
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Количество попыток миграции
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		// Пытаемся выполнить миграции
		err := storage.ExecuteMigrations(ctx, "/root/migrations")
		if err == nil {
			log.Println("Migrations applied successfully.")
			break // Если миграции прошли успешно, выходим из цикла
		}

		// Если произошла ошибка, выводим ее и пробуем снова
		log.Printf("Migration attempt %d failed: %v. Retrying...", i+1, err)

		// Ожидаем перед повторной попыткой
		time.Sleep(5 * time.Second)
	}

	// Если миграции не были выполнены после максимального количества попыток, завершаем выполнение
	if err != nil {
		log.Fatalf("Failed to apply migrations after %d attempts: %v", maxRetries, err)
	}

	logService := service.NewLogService(storage, *cfg)

	logApi := handler.NewLogStorageApi(logService)

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.JWTAuthInterceptor()),
	)
	pb.RegisterLogStorageServiceServer(grpcServer, logApi)

	log.Println("LogStorageService is running on port: ", cfg.Server.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
