package methods

import (
	"context"
	pg "crud/common-libs/postgres"
	shared "crud/common-libs/shared"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func create(ctx context.Context, note shared.Note) (string, error) {
	title, content, userId := note.Title, note.Content, note.UserId

	conn, err := pg.Pool.Acquire(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Release() // Release the connection back to the pool

	query, args, err := sq.Select("id").
		From("public.users").
		Where(sq.Eq{"id": userId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return "", err
	}

	var uid int
	if err := conn.QueryRow(ctx, query, args...).Scan(&uid); err != nil {
		fmt.Println("Error scanning UID:", err)
		return "", err
	}

	query, args, err = sq.Insert("public.notes").
		Columns("name", "text", "owner_id").
		Values(title, content, uid).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return "", err
	}

	var id string
	if err := conn.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		fmt.Println("Error scanning ID:", err)
		return "", err
	}

	return id, nil
}

func Create(ctx context.Context, data shared.CreateNoteData) (string, error) {
	note := data.Note

	id, err := create(ctx, note)
	if err != nil {
		return "", err
	}

	return id, nil
}
