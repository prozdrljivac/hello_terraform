package postgresdb

import (
	"context"
	"fmt"

	"github.com/prozdrljivac/hello_terraform/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("could not connect to DB: %w", err)
	}
	return pool, nil
}
