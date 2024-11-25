package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Структура для хранения информации о пользователе, извлеченной из JWT
type UserClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type SensorClaims struct {
	sensor_id string
	jwt.RegisteredClaims
}

func GenerateAPIJWT(secretKey string, sensorID string, expirationTime time.Time) (string, error) {
	claims := SensorClaims{
		sensor_id: sensorID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Указываем время истечения токена
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

// GenerateUserJWT генерирует JWT токен для пользователя.
func GenerateUserJWT(secretKey string, userID int64, username string, role string, expirationTime time.Time) (string, error) {
	// Создаем данные для токена
	claims := UserClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Указываем время истечения токена
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // Указываем время создания токена
		},
	}

	// Создаем новый токен с HMAC-SHA256 алгоритмом
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с использованием секретного ключа
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateUserJWT(tokenStr string, secretKey string) (bool, string, error) {
	// Парсим токен с использованием секретного ключа
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи соответствует HS256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		// Возвращаем секретный ключ в виде []byte для проверки подписи
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, "", fmt.Errorf("invalid token: %v", err)
	}

	// Проверяем, что токен действителен
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		// Проверка времени истечения срока действия токена (если есть)
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now().Add(-time.Second)) { // Даем немного времени на синхронизацию
			return false, "", fmt.Errorf("token has expired")
		}
		// Возвращаем роль пользователя и успешную валидацию
		return true, claims.Role, nil
	}

	return false, "", fmt.Errorf("invalid token claims")
}

func ValidateAPIJWT(tokenStr string, secretKey string) (bool, string, error) {
	// Парсим токен с использованием секретного ключа
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи соответствует HS256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		// Возвращаем секретный ключ в виде []byte для проверки подписи
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, "", fmt.Errorf("invalid token: %v", err)
	}

	// Проверяем, что токен действителен
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		// Проверка времени истечения срока действия токена (если есть)
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now().Add(-time.Second)) { // Даем немного времени на синхронизацию
			return false, "", fmt.Errorf("token has expired")
		}
		// Возвращаем информацию о токене и успешную валидацию
		return true, claims.Subject, nil
	}

	return false, "", fmt.Errorf("invalid token claims")
}
