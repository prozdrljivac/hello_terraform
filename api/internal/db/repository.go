package db

import (
	"context"

	"api/hello_terraform/internal/model"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, text string) (model.Message, error)
	ListMessages(ctx context.Context) ([]model.Message, error)
}
