package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	handler "github.com/imyazip/GoSIEM/auth/internal/server"
	auth "github.com/imyazip/GoSIEM/auth/internal/service"
	"github.com/imyazip/GoSIEM/auth/internal/storage"

	"github.com/imyazip/GoSIEM/auth/pkg/config"
	pb "github.com/imyazip/GoSIEM/auth/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	authService := *auth.NewAuthService(*storage, cfg)

	authHandler := handler.NewAuthAPI(&authService)

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err) // Логируем ошибку, если не удается слушать порт
	}

	// Создаем gRPC сервер
	grpcServer := grpc.NewServer()

	// Регистрируем наш сервис на сервере
	pb.RegisterAuthServiceServer(grpcServer, authHandler)
	reflection.Register(grpcServer)
	// Запускаем сервер
	log.Printf("Server is running on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err) // Логируем ошибку, если не удается запустить сервер
	}
}
