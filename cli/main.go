package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/imyazip/GoSIEM/cli/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	nginxAuthURL       = "localhost:80" // URL для AuthService через NGINX
	nginxLogStorageURL = "localhost:80" // URL для LogStorageService через NGINX
	username           = "admin"        // Пример логина
	password           = "admin"        // Пример пароля
)

var token string // Глобальный токен для авторизации

func main() {
	var err error

	// Ввод логина и пароля при запуске
	fmt.Print("Введите логин: ")
	var username string
	fmt.Scan(&username)

	fmt.Print("Введите пароль: ")
	var password string
	fmt.Scan(&password)

	// Авторизация и получение JWT
	token, err = loginAndGetJWT(username, password)
	if err != nil {
		log.Fatalf("Ошибка авторизации: %v", err)
	}

	log.Printf("Успешная авторизация. JWT: %s", token)

	// Отображение меню
	for {
		showMenu()
	}
}

// loginAndGetJWT выполняет запрос к AuthService для получения JWT токена
func loginAndGetJWT(username, password string) (string, error) {
	conn, err := grpc.Dial(nginxAuthURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", fmt.Errorf("не удалось подключиться к AuthService: %w", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	req := &pb.LoginRequest{
		Username: username,
		Password: password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.Login(ctx, req)
	if err != nil {
		return "", fmt.Errorf("не удалось выполнить вход: %w", err)
	}

	return resp.Token, nil
}

// showMenu отображает текстовое меню и обрабатывает выбор пользователя
func showMenu() {
	fmt.Println("\n--- Главное меню ---")
	fmt.Println("1. Управление пользователями")
	fmt.Println("2. Управление правилами")
	fmt.Println("3. Просмотр событий")
	fmt.Println("4. Просмотр событий безопасности")
	fmt.Println("0. Выход")

	fmt.Print("Введите ваш выбор: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		manageUsers()
	case 2:
		manageRules()
	case 3:
		viewLogs()
	case 4:
		viewSecurityEvents()
	case 0:
		fmt.Println("Выход из приложения.")
		exitApplication()
	default:
		fmt.Println("Неверный выбор. Попробуйте снова.")
	}
}

// manageUsers обрабатывает управление пользователями
func manageUsers() {
	fmt.Println("\n--- Управление пользователями ---")
	fmt.Println("1. Создать пользователя")
	fmt.Println("2. Удалить пользователя")
	fmt.Println("0. Назад в меню")

	fmt.Print("Введите ваш выбор: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		createUser()
	case 2:
		deleteUser()
	case 0:
		return
	default:
		fmt.Println("Неверный выбор. Попробуйте снова.")
	}
}

func createUser() {
	fmt.Print("Введите имя пользователя: ")
	var username string
	fmt.Scan(&username)

	fmt.Print("Введите пароль: ")
	var password string
	fmt.Scan(&password)

	fmt.Print("Введите роль (1 - admin, 2 - analyst, 3 - viewer): ")
	var role int
	fmt.Scan(&role)

	conn, err := grpc.Dial(nginxAuthURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+token)

	req := &pb.CreateUserRequest{
		Username: username,
		Password: password,
		Role:     int64(role),
	}

	resp, err := client.CreateUser(ctx, req)
	if err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
		return
	}

	if resp.Success {
		fmt.Println("Пользователь успешно создан.")
	} else {
		fmt.Println("Не удалось создать пользователя.")
	}
}

func deleteUser() {
	fmt.Print("Введите имя пользователя для удаления: ")
	var username string
	fmt.Scan(&username)

	conn, err := grpc.Dial(nginxAuthURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+token)

	req := &pb.DeleteUserRequest{Username: username}

	resp, err := client.DeleteUser(ctx, req)
	if err != nil {
		fmt.Printf("Failed to delete user: %v\n", err)
		return
	}

	if resp.Success {
		fmt.Println("Пользователь успешно удален.")
	} else {
		fmt.Println("Не удалось удалить пользователя.")
	}
}

// manageRules обрабатывает управление правилами
func manageRules() {
	fmt.Println("\n--- Управление правилами ---")
	fmt.Println("1. Добавить правило")
	fmt.Println("2. Просмотреть правила")
	fmt.Println("0. Назад в меню")

	fmt.Print("Введите ваш выбор: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		addRule()
	case 2:
		getRules()
	case 0:
		return
	default:
		fmt.Println("Неверный выбор. Попробуйте снова.")
	}
}

func addRule() {
	fmt.Print("Введите новое правило (JSON): ")
	var rule string
	fmt.Scan(&rule)

	conn, err := grpc.Dial(nginxLogStorageURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewLogStorageServiceClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+token)

	req := &pb.AddRuleRequest{Rule: rule}

	resp, err := client.AddRule(ctx, req)
	if err != nil {
		fmt.Printf("Failed to add rule: %v\n", err)
		return
	}

	fmt.Printf("Правило успешно добавлено с ID: %d\n", resp.RuleId)
}

func getRules() {
	conn, err := grpc.Dial(nginxLogStorageURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewLogStorageServiceClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+token)

	req := &pb.GetRulesRequest{}

	resp, err := client.GetRules(ctx, req)
	if err != nil {
		fmt.Printf("Failed to get rules: %v\n", err)
		return
	}

	fmt.Println("Правила:")
	for _, rule := range resp.Rules {
		fmt.Printf("ID: %d, Rule: %s\n", rule.Id, rule.Rule)
	}
}

// viewLogs и viewSecurityEvents – упрощенные примеры
func viewLogs() {
	fmt.Println("\n--- Просмотр событий ---")
	// Реализация аналогична другим методам
}

func viewSecurityEvents() {
	fmt.Println("\n--- Просмотр событий безопасности ---")

	fmt.Print("Введите лимит событий для отображения: ")
	var limit int
	fmt.Scan(&limit)

	fmt.Print("Показать только непрочитанные события? (1 - Да, 0 - Нет): ")
	var onlyUnreadInput int
	fmt.Scan(&onlyUnreadInput)
	onlyUnread := onlyUnreadInput == 1

	conn, err := grpc.Dial(nginxLogStorageURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Ошибка подключения: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewLogStorageServiceClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+token)

	req := &pb.GetSecurityEventsRequest{
		Limit:      int32(limit),
		OnlyUnread: onlyUnread,
	}

	resp, err := client.GetSecurityEvents(ctx, req)
	if err != nil {
		fmt.Printf("Ошибка получения событий безопасности: %v\n", err)
		return
	}

	fmt.Println("\nСобытия безопасности:")
	if len(resp.Events) == 0 {
		fmt.Println("Нет событий для отображения.")
		return
	}

	for _, event := range resp.Events {
		fmt.Printf("ID: %d\n", event.Id)
		fmt.Printf("Тип: %s\n", event.EventType)
		fmt.Printf("Описание: %s\n", event.EventDescription)
		fmt.Printf("Обнаружено: %s\n", event.DetectedAt.AsTime().Format(time.RFC1123))
		fmt.Printf("Прочитано: %t\n", event.ReadFlag)
		fmt.Println("---")
	}
}

func exitApplication() {
	fmt.Println("Спасибо за использование программы!")
	time.Sleep(time.Second)
	log.Fatal("Приложение завершено.")
}
