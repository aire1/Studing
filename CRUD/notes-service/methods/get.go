package methods

import (
	"context"
	pg "crud/common-libs/postgres"
	shared "crud/common-libs/shared"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func get(ctx context.Context, offset, count int) (*[]shared.Note, error) {
	if offset < 0 {
		offset = 0
	}
	if count < 0 {
		count = 100
	}

	conn, err := pg.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	username := ctx.Value(shared.UsernameKey).(string)

	query, args, err := sq.Select("id").
		From("public.users").
		Where(sq.Eq{"username": username}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var uid int
	if err := conn.QueryRow(ctx, query, args...).Scan(&uid); err != nil {
		fmt.Println("Error scanning UID:", err)
		return nil, err
	}

	query, args, err = sq.Select("id, name, text, created_at::text").
		From("public.notes").
		Where(
			sq.Eq{"owner_id": uid},
			sq.NotEq{"deleted": true}).
		Offset(uint64(offset)).
		Limit(uint64(count)).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	res, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var notes = []shared.Note{}
	var note = shared.Note{}
	for res.Next() {
		note = shared.Note{}
		if err := res.Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return &notes, nil
}

func Get(ctx context.Context, data shared.GetNoteData) (*[]shared.Note, error) {
	ctx = context.WithValue(ctx, shared.UsernameKey, data.Login)

	notes, err := get(ctx, data.Offset, data.Count)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
