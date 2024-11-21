// Package server представляет собой реализацию rpc методов описанных в .proto файле
package server

import (
	"context"

	auth "github.com/imyazip/GoSIEM/auth/internal/service"
	pb "github.com/imyazip/GoSIEM/auth/proto"
)

// Стркутура, реализующая функционал API
type AuthAPI struct {
	pb.UnimplementedAuthServiceServer
	service *auth.AuthService
}

// Конструктор AuthHandler
func NewAuthAPI(service *auth.AuthService) *AuthAPI {
	return &AuthAPI{service: service}
}

// Реализация rpc метода GenerateJWTForSensor, генерирует JWT токен для сенора, используя переданный API-ключ.
func (h *AuthAPI) GenerateJWTForSensor(ctx context.Context, req *pb.GenerateJWTForSensorRequest) (*pb.GenerateJWTForSensorResponse, error) {
	token, err := h.service.GenerateJWTFromAPIKey(ctx, req.ApiKey)

	if err != nil {
		return nil, err
	}

	return &pb.GenerateJWTForSensorResponse{
		Token: token,
	}, nil
}
