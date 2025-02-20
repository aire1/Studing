package check_authorization

import (
	"context"
	pb "crud/auth-service/proto"
	rd "crud/common-libs/redis"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

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

func CheckAuthorization(ctx *context.Context, req *pb.AuthCheckRequest) error {
	var jwtErr, redisErr error
	var res *jwt.Claims
	var redisValid bool

	res, jwtErr = jwt.ValidateToken(req.JwtToken)
	if jwtErr != nil {
		return errors.Errorf("error validating token")
	}

	redisValid, redisErr = ValidateJWTFromRedis(res.Username, req.JwtToken, *ctx)

	if redisErr != nil {
		return errors.Errorf("error validating token")
	} else if res == nil {
		return errors.Errorf("jwt not valid")
	} else if errors.Is(redisErr, redis.Nil) {
		return errors.Errorf("can't find jwt session: %v", redisErr)
	} else if !redisValid {
		return errors.Errorf("client token != stored token")
	}

	*ctx = context.WithValue(*ctx, "username", res.Username)
	*ctx = context.WithValue(*ctx, "uid", res.Uid)

	return nil
}
