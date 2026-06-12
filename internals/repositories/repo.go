package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool is the shared PostgreSQL connection pool used by every repository.
var Pool *pgxpool.Pool

// Init opens the connection pool against the given DATABASE_URL.
func Init(databaseURL string) error {
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return err
	}
	if err := pool.Ping(context.Background()); err != nil {
		return err
	}
	Pool = pool
	return nil
}
