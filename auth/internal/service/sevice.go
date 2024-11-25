// Package service выполняет основную бизнес-логику API запросов сервиса аутендификации и авторизации.
package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/imyazip/GoSIEM/auth/internal/models"
	"github.com/imyazip/GoSIEM/auth/internal/storage"
	"github.com/imyazip/GoSIEM/auth/pkg/config"
	"github.com/imyazip/GoSIEM/auth/pkg/hash"
	"github.com/imyazip/GoSIEM/auth/pkg/jwt"
	"google.golang.org/grpc/metadata"
)

// AuthService представляет собой основной сервис для аутентификации и авторизации.
type AuthService struct {
	db     storage.Storage
	config *config.Config
}

// NewAuthService создает и возвращает новый экземпляр AuthService.
// db - слой доступа к данным.
// config - конфигурация сервиса.
func NewAuthService(db storage.Storage, config *config.Config) *AuthService {
	return &AuthService{db: db, config: config}
}

// GenerateJWTFromAPIKey генерирует JWT токен для переданного API ключа.
// Он проверяет, существует ли API ключ в базе данных, и если ключ действителен,
// генерирует и возвращает JWT токен с заданным временем истечения.
func (s *AuthService) GenerateJWTFromAPIKey(ctx context.Context, apiKey string, sensor models.Sensor) (string, error) {
	exists, err := s.db.CheckSensorExists(ctx, sensor.Sensor_id)
	if err != nil {
		log.Printf("Error finding sensor in db: %v", err)
	}

	if !exists {
		err = s.db.InsertSensor(ctx, sensor)
		if err != nil {
			log.Printf("Error inserting sensor in db: %v", err)
		}
	}

	err = s.db.UpdateSensor(ctx, sensor)
	if err != nil {
		log.Printf("Error updating sensor in db: %v", err)
	}

	expirationTime := time.Now().Add(time.Duration(s.config.JWT.ExpirationMinutes) * time.Minute)
	valid, err := s.db.FindAPIKeyInStorage(ctx, apiKey)
	if err != nil {
		log.Printf("Error validating API key: %v", err)
		return "", errors.New("failed to validate API key")
	}
	if !valid {
		log.Printf("Invalid API key: %s", apiKey)
		return "", errors.New("invalid API key")
	}

	return jwt.GenerateAPIJWT(s.config.JWT.SecretKey, sensor.Sensor_id, expirationTime)
}

func (s *AuthService) CreateNewUser(ctx context.Context, username string, password string, roleID int64) error {
	// Извлекаем метаданные из контекста для проверки jwt токена
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("failed to get metadata")
		return errors.New("failed to get metadata")
	}

	// Извлекаем токен из метаданных
	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		log.Printf("missing authorization token")
		return errors.New("missing authorization token")
	}

	// Токен должен быть в формате "Bearer <token>"
	token := authHeader[0][7:] // Убираем "Bearer " из строки

	// Проверяем валидность токена
	valid, role, err := jwt.ValidateUserJWT(token, s.config.JWT.SecretKey)
	if err != nil || !valid || role != "admin" {
		fmt.Println(role)
		return errors.New("invalid or insufficient token")
	}

	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return errors.New("failed hashing password")
	}

	err = s.db.InsertUser(ctx, username, hashedPassword, roleID)
	if err != nil {
		return errors.New("failed inserting user to database")
	}

	log.Printf("Added user %s with roleID %d to database", username, roleID)
	return nil

}

func (s *AuthService) GenerateJWTFromUser(ctx context.Context, username string, password string) (string, error) {
	user, err := s.db.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	err = hash.ComparePasswordHash(user.Password, password)
	if err != nil {
		return "", err
	}

	role, err := s.db.GetRoleNameByID(ctx, user.RoleID)
	if err != nil {
		return "", err
	}
	fmt.Println(s.config.JWT.ExpirationMinutes)
	expirationTime := time.Now().Add(time.Duration(s.config.JWT.ExpirationMinutes) * time.Minute)
	fmt.Println(expirationTime)
	token, err := jwt.GenerateUserJWT(s.config.JWT.SecretKey, user.ID, user.Username, role, expirationTime)
	if err != nil {
		return "", err
	}
	return token, nil

}
