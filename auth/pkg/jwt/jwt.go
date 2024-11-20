package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(secretKey string, subject string, expirationTime time.Time) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
