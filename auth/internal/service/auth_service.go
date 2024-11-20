package auth

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/imyazip/GoSIEM/auth/internal/storage"
	"github.com/imyazip/GoSIEM/auth/pkg/config"
	"github.com/imyazip/GoSIEM/auth/pkg/jwt"
)

type AuthService struct {
	db     storage.Storage
	config *config.Config
}

func NewAuthService(db storage.Storage, config *config.Config) *AuthService {
	return &AuthService{db: db, config: config}
}

func (s *AuthService) GenerateJWT(ctx context.Context, apiKey string) (string, error) {
	valid, err := s.db.ValidateAPIKey(ctx, apiKey)
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
