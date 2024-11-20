package handler

import (
	"context"

	auth "github.com/imyazip/GoSIEM/auth/internal/service"
	pb "github.com/imyazip/GoSIEM/auth/proto"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service *auth.AuthService
}

func NewAuthHandler(service *auth.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) ValidateKey(ctx context.Context, req *pb.ValidateKeyRequest) (*pb.ValidateKeyResponse, error) {
	token, err := h.service.GenerateJWT(ctx, req.Key)

	if err != nil {
		return nil, err
	}

	return &pb.ValidateKeyResponse{
		Token: token,
	}, nil
}
