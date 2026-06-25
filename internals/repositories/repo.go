package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

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
