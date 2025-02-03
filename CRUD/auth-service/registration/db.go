package registration

import (
	"context"
	"fmt"

	pg "crud/auth-service/common-libs/postgres"
)

func RegisterClient(username string, passhash string) error {
	conn, err := pg.Pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		fmt.Sprintf(`insert into public.users (username, password_hash) values ('%s', '%s');`, username, passhash)).Scan()
	if err != nil {
		return err
	}

	return nil
}
