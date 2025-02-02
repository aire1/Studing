package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("studing") //знаю, нужно переность в параметры окружения, но это учебный проект)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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
