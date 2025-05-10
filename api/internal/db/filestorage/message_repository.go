package filestorage

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/prozdrljivac/hello_terraform/internal/model"
)

type FileStorageMessageRepository struct {
	file *os.File
}

func NewFileStorageMessageRepository(file *os.File) *FileStorageMessageRepository {
	return &FileStorageMessageRepository{file: file}
}

func (r *FileStorageMessageRepository) CreateMessage(
	ctx context.Context, text string,
) (model.Message, error) {

	_, _ = r.file.Seek(0, io.SeekStart)
	var msgs []model.Message
	if err := json.NewDecoder(r.file).Decode(&msgs); err != nil && !errors.Is(err, io.EOF) {
		return model.Message{}, err
	}

	nextID := 1
	if len(msgs) > 0 {
		nextID = msgs[len(msgs)-1].ID + 1
	}
	msg := model.Message{ID: nextID, Text: text}
	msgs = append(msgs, msg)

	if err := r.file.Truncate(0); err != nil {
		return model.Message{}, err
	}
	_, _ = r.file.Seek(0, io.SeekStart)
	if err := json.NewEncoder(r.file).Encode(msgs); err != nil {
		return model.Message{}, err
	}

	return msg, nil
}

func (r *FileStorageMessageRepository) ListMessages(
	ctx context.Context,
) ([]model.Message, error) {

	if _, err := r.file.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	var msgs []model.Message
	err := json.NewDecoder(r.file).Decode(&msgs)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	return msgs, nil
}
