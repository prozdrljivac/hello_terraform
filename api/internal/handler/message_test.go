package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"api/hello_terraform/internal/handler"
	"api/hello_terraform/internal/model"
)

type mockRepo struct {
	ListFunc func(ctx context.Context) ([]model.Message, error)
}

func (m *mockRepo) ListMessages(ctx context.Context) ([]model.Message, error) {
	return m.ListFunc(ctx)
}

func (m *mockRepo) CreateMessage(ctx context.Context, text string) (model.Message, error) {
	return model.Message{}, nil
}

func TestListMessages_ReturnsStatusOK(t *testing.T) {
	repo := &mockRepo{
		ListFunc: func(ctx context.Context) ([]model.Message, error) {
			return []model.Message{}, nil
		},
	}
	h := handler.NewMessageHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestListMessages_ReturnsMessagesInJSON(t *testing.T) {
	repo := &mockRepo{
		ListFunc: func(ctx context.Context) ([]model.Message, error) {
			return []model.Message{
				{ID: 1, Text: "Hello"},
				{ID: 2, Text: "World"},
			}, nil
		},
	}
	h := handler.NewMessageHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var result []model.Message
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if len(result) != 2 || result[0].ID != 1 || result[0].Text != "Hello" || result[1].ID != 2 || result[1].Text != "World" {
		t.Errorf("unexpected response data: %+v", result)
	}
}
