package methods

import (
	"context"
	pg "crud/common-libs/postgres"
	shared "crud/common-libs/shared"

	sq "github.com/Masterminds/squirrel"
)

func Create(ctx context.Context, note shared.Note) (string, error) {
	title, content, userId := note.Title, note.Content, note.UserId

	conn, err := pg.Pool.Acquire(ctx)
	if err != nil {
		return "", err
	}

	query, args, err := sq.Insert("public.notes").
		Columns("name", "text", "user_id").
		Values(title, content, userId).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return "", err
	}

	var id string
	if err := conn.QueryRow(ctx, query, args).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
