// Package server представляет собой реализацию rpc методов описанных в .proto файле
package server

import (
	"context"

	"github.com/imyazip/GoSIEM/auth/internal/models"
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
	var sensor models.Sensor
	sensor.Sensor_id = req.SensorId
	sensor.Name = req.Name
	sensor.Hostname = req.Hostname
	sensor.Os_version = req.OsVersion
	sensor.Sensor_type = req.SensorType
	sensor.Agent_version = req.AgentVersion

	token, err := h.service.GenerateJWTFromAPIKey(ctx, req.ApiKey, sensor)

	if err != nil {
		return nil, err
	}

	return &pb.GenerateJWTForSensorResponse{
		Token: token,
	}, nil
}

func (h *AuthAPI) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := h.service.CreateNewUser(ctx, req.Username, req.Password, req.Role)
	if err != nil {
		println(err)
		return &pb.CreateUserResponse{
			Success: false,
		}, err
	}

	return &pb.CreateUserResponse{
		Success: true,
	}, nil

}

func (h *AuthAPI) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := h.service.GenerateJWTFromUser(ctx, req.Username, req.Password)

	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil

}
