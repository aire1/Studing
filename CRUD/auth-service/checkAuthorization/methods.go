package check_authorization

import (
	"context"
	rd "crud/common-libs/redis"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	shared "crud/common-libs/shared"
	jwt "crud/common-libs/shared/jwt"
)

// Функция для проверки JWT идентичности Redis
func ValidateJWTFromRedis(userID, token string, ctx context.Context) (bool, error) {
	key := "session:" + userID
	storedToken, err := rd.Client.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return storedToken == token, nil
}

func CheckAuthorization(ctx context.Context, data shared.AuthorizationCheckData) error {
	var wg sync.WaitGroup
	var jwtErr, redisErr error
	var res *jwt.Claims
	var redisValid bool

	wg.Add(2)

	go func() {
		defer wg.Done()
		res, jwtErr = jwt.ValidateToken(data.JwtToken)
	}()

	go func() {
		defer wg.Done()
		redisValid, redisErr = ValidateJWTFromRedis(data.Username, data.JwtToken, ctx)
	}()

	wg.Wait()

	if jwtErr != nil {
		return errors.Errorf("error validating token")
	}
	if res == nil {
		return errors.Errorf("jwt not valid")
	}
	if redisErr == redis.Nil {
		return errors.Errorf("can't find jwt session: %v", redisErr)
	}
	if redisErr != nil {
		return errors.Errorf("error validating token from redis: %v", redisErr)
	}
	if !redisValid {
		return errors.Errorf("client token != stored token")
	}

	return nil
}
