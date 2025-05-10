package postgresdb

import (
	"context"

	"github.com/prozdrljivac/hello_terraform/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresMessageRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresMessageRepository(pool *pgxpool.Pool) *PostgresMessageRepository {
	return &PostgresMessageRepository{pool: pool}
}

func (r *PostgresMessageRepository) CreateMessage(ctx context.Context, text string) (model.Message, error) {
	var msg model.Message
	err := r.pool.QueryRow(ctx, `
		INSERT INTO messages (text) VALUES ($1)
		RETURNING id, text;
	`, text).Scan(&msg.ID, &msg.Text)
	return msg, err
}

func (r *PostgresMessageRepository) ListMessages(ctx context.Context) ([]model.Message, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, text FROM messages ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.ID, &msg.Text); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
