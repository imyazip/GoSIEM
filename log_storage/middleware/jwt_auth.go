package middleware

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// JWTAuthInterceptor извлекает JWT из метаданных и добавляет его в контекст
func JWTAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("missing metadata")
		}

		jwt := md["authorization"]
		if len(jwt) == 0 {
			return nil, fmt.Errorf("missing JWT token")
		}

		newCtx := context.WithValue(ctx, "jwt", jwt[0])
		return handler(newCtx, req)
	}
}
