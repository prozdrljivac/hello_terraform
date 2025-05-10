package filestorage

import (
	"context"
	"fmt"
	"os"
)

func NewFileStorage(ctx context.Context, path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("could not create the storage: %w", err)
	}

	return file, nil
}
