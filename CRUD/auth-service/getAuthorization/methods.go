package get_authorizatrion

import (
	"context"
	"time"

	"github.com/pkg/errors"

	rd "crud/common-libs/redis"
	shared "crud/common-libs/shared"
	jwt "crud/common-libs/shared/jwt"
)

func Authorize(ctx context.Context, data shared.AuthorizationGetData) (string, error) {
	passhash, err := GetUserPasshash(data.Login)
	if err != nil {
		return "", errors.Errorf("can't get user: %v", err)
	} else if passhash == "" {
		return "", errors.Errorf("user not exists")
	} else if passhash != data.Passhash {
		return "", errors.Errorf("invalid credentials")
	}

	uid, err := GetUserId(data.Login)
	if err != nil {
		return "", errors.Errorf("can't get user id: %v", err)
	}
	token, err := jwt.GenerateToken(data.Login, uid, time.Hour*24)
	if err != nil {
		return "", errors.Errorf("can't generate jwt: %v", err)
	}

	if err = jwt.StoreTokenInRedis(ctx, rd.Client, data.Login, token); err != nil {
		return "", errors.Errorf("can't push jwt token into reddis: %v", err)
	}

	return token, nil
}
