package methods

import (
	"context"
	pg "crud/common-libs/postgres"
	shared "crud/common-libs/shared"
	"fmt"

	logger "crud/notes-service/logger"

	sq "github.com/Masterminds/squirrel"
)

// Удаление заметки, либо пометка о удалении
func update(ctx context.Context, note shared.Note) error {
	conn, err := pg.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	query, args, err := sq.Update("public.notes").
		Set("text", note.Content).
		Set("name", note.Title).
		Where(sq.Eq{"id": note.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	if err := conn.QueryRow(ctx, query, args...).Scan(); err != nil {
		logger.LogWithCaller(fmt.Sprintf("Error executing SQL: %v", err))
		return err
	}

	return nil
}

func Update(ctx context.Context, data shared.UpdateNoteData) error {
	note := data.Note

	err := update(ctx, note)
	if err != nil {
		return err
	}

	return nil
}
