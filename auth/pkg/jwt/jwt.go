package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Структура для хранения информации о пользователе, извлеченной из JWT
type UserClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(secretKey string, subject string, expirationTime time.Time) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func ValidateUserJWT(tokenStr string, secretKey string) (bool, string, error) {
	// Парсим токен с использованием секретного ключа
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи соответствует HS256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		// Возвращаем секретный ключ для проверки подписи
		return secretKey, nil
	})
	if err != nil {
		return false, "", fmt.Errorf("invalid token: %v", err)
	}
	// Проверяем, что токен действителен
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		// Проверка времени истечения срока действия токена (если есть)
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return false, "", fmt.Errorf("token has expired")
		}
		// Возвращаем роль пользователя и успешную валидацию
		return true, claims.Role, nil
	}

	return false, "", fmt.Errorf("invalid token claims")
}
