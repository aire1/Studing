package get_authorizatrion

import (
	"context"
	"time"

	rd "crud/common-libs/redis"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	shared "crud/common-libs/shared"
)

// Secret key для подписи JWT
var secretKey = []byte("super-secret-key")

var ctx = context.Background()

// Функция для генерации JWT
func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	return token.SignedString(secretKey)
}

// Функция для сохранения JWT в Redis
func StoreTokenInRedis(userID, token string) error {
	key := "session:" + userID
	err := rd.Client.Set(ctx, key, token, time.Hour).Err()
	return err
}

func Authorize(ctx context.Context, data shared.AuthorizationGetData) error {
	passhash, err := GetUserPasshash(data.Login)
	if err != nil {
		return errors.Errorf("can't get user: %v", err)
	} else if passhash == "" {
		return errors.Errorf("user not exists")
	} else if passhash != data.Passhash {
		return errors.Errorf("invalid credentials")
	}

	token, err := GenerateJWT(data.Login)
	if err != nil {
		return errors.Errorf("can't generate jwt: %v", err)
	}

	if err = StoreTokenInRedis(data.Login, token); err != nil {
		return errors.Errorf("can't push jwt token into reddis: %v", err)
	}

	return nil
}
