package methods

import (
	"context"
	pg "crud/common-libs/postgres"
	shared "crud/common-libs/shared"

	sq "github.com/Masterminds/squirrel"
)

func get(ctx context.Context, userId, offset, count int) (*[]shared.Note, error) {
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

	query, args, err := sq.Select("id, name, text, created_at").
		From("public.notes").
		Where(
			sq.Eq{"user_id": userId},
			sq.NotEq{"deleted": true}).
		Offset(uint64(offset)).
		Limit(uint64(count)).ToSql()

	res, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var notes = []shared.Note{}
	var note = shared.Note{}
	for res.Next() {
		note = shared.Note{}
		if err := res.Scan(&note.Id, &note.Title, &note.Content, &note.CreationDate); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return &notes, nil
}
