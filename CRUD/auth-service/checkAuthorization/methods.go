package check_authorization

import (
	"context"
	rd "crud/auth-service/common-libs/redis"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// Функция для проверки JWT из Redis
func ValidateJWTFromRedis(userID, token string, ctx context.Context) (bool, error) {
	key := "session:" + userID
	storedToken, err := rd.Client.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return storedToken == token, nil
}

type AuthData struct {
	Username string `json:"username"`
	JwtToken string `json:"jwtToken"`
	Taskid   string `json:"taskid"`
}

func CheckAuthorization(ctx context.Context, data AuthData) (bool, error) {
	if res, err := ValidateJWTFromRedis(data.Username, data.JwtToken, ctx); err == redis.Nil {
		return false, errors.Errorf("can't find jwt in redis: %v", err)
	} else if err != nil {
		return false, errors.Errorf("error validating token from redis: %v", err)
	} else {
		return res, nil
	}
}
