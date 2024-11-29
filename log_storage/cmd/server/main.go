package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/go-sql-driver/mysql"
	handler "github.com/imyazip/GoSIEM/log-storage/internal/server"
	"github.com/imyazip/GoSIEM/log-storage/internal/service"
	"github.com/imyazip/GoSIEM/log-storage/internal/storage"
	"github.com/imyazip/GoSIEM/log-storage/pkg/config"
	pb "github.com/imyazip/GoSIEM/log-storage/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	time.Sleep(15 * time.Second)
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

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err) // Логируем ошибку, если не удается слушать порт
	}

	// Создаем gRPC сервер
	grpcServer := grpc.NewServer()

	// Регистрируем наш сервис на сервере
	pb.RegisterLogStorageServiceServer(grpcServer, logApi)
	reflection.Register(grpcServer)
	// Запускаем сервер
	log.Printf("Server is running on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err) // Логируем ошибку, если не удается запустить сервер
	}
}
