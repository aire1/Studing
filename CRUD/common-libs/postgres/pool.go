package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Pool *pgxpool.Pool
	err  error
)

func CreatePool() {
	Pool, err = pgxpool.New(context.Background(), "postgresql://postgres:123@localhost:5433/postgres")
	if err != nil {
		log.Fatalf("unable to connect to database: %s", err.Error())
	}
}
