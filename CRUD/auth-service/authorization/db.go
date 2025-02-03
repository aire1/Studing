package authorizatrion

import (
	"context"
	"fmt"

	pg "crud/auth-service/common-libs/postgres"
)

func GetUserPasshash(username string) (string, error) {
	conn, err := pg.Pool.Acquire(context.Background())
	if err != nil {
		return "", err
	}
	defer conn.Release()

	var passhash string
	err = conn.QueryRow(
		context.Background(),
		fmt.Sprintf(`select password_hash from public.users where username = '%s';`, username),
	).Scan(&passhash)
	if err != nil {
		return "", err
	}

	return passhash, nil
}
