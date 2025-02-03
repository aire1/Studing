package authorizatrion

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

var secretKey = []byte("studing")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	// Определяем срок действия токена
	expirationTime := time.Now().Add(24 * time.Hour) // 24 часа

	// Создаем claims (данные, которые будут в токене)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Время истечения токена
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // Время создания токена
			Subject:   "auth",                             // Назначение токена
		},
	}

	// Создаем токен с подписью
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Функция для проверки JWT-токена
func ValidateToken(tokenString string) (*Claims, error) {
	// Парсим токен
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи - HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Проверяем валидность токена и извлекаем claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

type AuthData struct {
	Login    string `json:"login"`
	Passhash string `json:"passhash"`
	Taskid   string `json:"taskid"`
}

func Authorize(ctx context.Context, data AuthData) (string, error) {
	passhash, err := GetUserPasshash(data.Login)
	if err != nil {
		return "", errors.Errorf("can't get user: %v", err)
	} else if passhash == "" {
		return "", errors.Errorf("user not exists")
	} else if passhash != data.Passhash {
		return "", errors.Errorf("invalid credentials")
	}

	return "", nil
}
