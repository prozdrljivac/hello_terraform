package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"api/hello_terraform/internal/handler"
	"api/hello_terraform/internal/model"
)

type mockRepo struct {
	ListFunc   func(ctx context.Context) ([]model.Message, error)
	CreateFunc func(ctx context.Context, text string) (model.Message, error)
}

func (m *mockRepo) ListMessages(ctx context.Context) ([]model.Message, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx)
	}
	return nil, nil
}

func (m *mockRepo) CreateMessage(ctx context.Context, text string) (model.Message, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, text)
	}
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

func TestCreateMessage_ReturnsStatusCreated(t *testing.T) {
	repo := &mockRepo{
		CreateFunc: func(ctx context.Context, text string) (model.Message, error) {
			return model.Message{ID: 1, Text: text}, nil
		},
	}
	h := handler.NewMessageHandler(repo)

	body := []byte(`{"text": "Hello"}`)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestCreateMessage_ReturnsCorrectJSON(t *testing.T) {
	repo := &mockRepo{
		CreateFunc: func(ctx context.Context, text string) (model.Message, error) {
			return model.Message{ID: 1, Text: text}, nil
		},
	}
	h := handler.NewMessageHandler(repo)

	body := []byte(`{"text": "Hello"}`)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	var msg model.Message
	if err := json.NewDecoder(w.Body).Decode(&msg); err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if msg.ID != 1 || msg.Text != "Hello" {
		t.Errorf("unexpected message response: %+v", msg)
	}
}

func TestCreateMessage_ReturnsBadRequestForMissingText(t *testing.T) {
	repo := &mockRepo{}
	h := handler.NewMessageHandler(repo)

	body := []byte(`{"no_text": "missing"}`)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}
