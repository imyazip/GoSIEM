// Package service выполняет основную бизнес-логику API запросов сервиса аутендификации и авторизации.
package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/imyazip/GoSIEM/auth/internal/storage"
	"github.com/imyazip/GoSIEM/auth/pkg/config"
	"github.com/imyazip/GoSIEM/auth/pkg/jwt"
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
func (s *AuthService) GenerateJWTFromAPIKey(ctx context.Context, apiKey string) (string, error) {
	valid, err := s.db.FindAPIKeyInStorage(ctx, apiKey)
	if err != nil {
		log.Printf("Error validating API key: %v", err)
		return "", errors.New("failed to validate API key")
	}
	if !valid {
		log.Printf("Invalid API key: %s", apiKey)
		return "", errors.New("invalid API key")
	}

	expirationTime := time.Now().Add(time.Duration(s.config.JWT.ExpirationMinutes) * time.Minute)
	return jwt.GenerateJWT(s.config.JWT.SecretKey, apiKey, expirationTime)
}
