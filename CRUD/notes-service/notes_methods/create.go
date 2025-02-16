package methods

import (
	"context"
	pg "crud/common-libs/postgres"
	shared "crud/common-libs/shared"
	"fmt"
)

func create(ctx context.Context, note shared.Note) error {
	title, content, userId := note.Title, note.Content, note.UserId

	conn, err := pg.Pool.Acquire(ctx)
	if err != nil {
		return err
	}

	if err := conn.QueryRow(ctx,
		fmt.Sprintf(`intert into public.notes (name, text, user_id) values ('%s', '%s', '%s');`, title, content, userId),
	).Scan(); err != nil {
		return err
	}

	return nil
}
