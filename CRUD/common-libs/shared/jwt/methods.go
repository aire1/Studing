package jwt

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
)

var jwtKey = []byte("studing") //знаю, нужно переность в параметры окружения, но это учебный проект)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Функция для сохранения JWT в Redis
func StoreTokenInRedis(ctx context.Context, client *redis.Client, userID string, token string) error {
	key := "session:" + userID
	err := client.Set(ctx, key, token, time.Hour).Err()
	return err
}

func GenerateToken(username string, expirationDate time.Duration) (string, error) {
	expiration := time.Now().Add(expirationDate)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
