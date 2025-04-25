package db

import (
	"context"

	"github.com/prozdrljivac/hello_terraform/internal/model"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, text string) (model.Message, error)
	ListMessages(ctx context.Context) ([]model.Message, error)
}
