package registration

import (
	"context"
	"fmt"

	pg "common-libs/postgres"
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
		fmt.Sprintf(`select password_hash from users where username = '%s'`, username),
	).Scan(&passhash)
	if err != nil {
		return "", err
	}

	return passhash, nil
}

func RegisterClient(username string, passhash string) error {
	conn, err := pg.Pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		fmt.Sprintf(`insert into users values (username, password_hash) ('%s', '%s')`, username, passhash)).Scan()
	if err != nil {
		return err
	}

	return nil
}
