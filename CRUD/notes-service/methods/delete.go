package methods

import (
	"context"
	pg "crud/common-libs/postgres"
	shared "crud/common-libs/shared"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// Удаление заметки, либо пометка о удалении
func delete(ctx context.Context, noteId int, force bool) error {
	conn, err := pg.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	var query string
	var args []interface{}
	if force {
		query, args, err = sq.Delete("public.notes").
			Where(sq.Eq{"id": noteId}).
			PlaceholderFormat(sq.Dollar).
			ToSql()
	} else {
		query, args, err = sq.Update("public.notes").
			Set("deleted", true).
			Where(sq.Eq{"id": noteId}).
			PlaceholderFormat(sq.Dollar).
			ToSql()
	}

	if err != nil {
		return err
	}

	if err := conn.QueryRow(ctx, query, args...).Scan(); err != nil {
		fmt.Println("Error executing SQL:", err)
		return err
	}

	return nil
}

func Delete(ctx context.Context, data shared.DeleteNoteData) error {
	noteId := data.NoteId

	err := delete(ctx, noteId, false)
	if err != nil {
		return err
	}

	return nil
}
