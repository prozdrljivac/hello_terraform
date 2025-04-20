package handler

import (
	"encoding/json"
	"net/http"

	"api/hello_terraform/internal/db"
)

type MessageHandler struct {
	repo db.MessageRepository
}

func NewMessageHandler(repo db.MessageRepository) *MessageHandler {
	return &MessageHandler{repo: repo}
}

func (h *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodGet:
		messages, err := h.repo.ListMessages(ctx)
		if err != nil {
			http.Error(w, "Failed to list messages", http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusOK, messages)

	case http.MethodPost:
		text := r.FormValue("text")
		if text == "" {
			http.Error(w, "Missing 'text' field", http.StatusBadRequest)
			return
		}
		msg, err := h.repo.CreateMessage(ctx, text)
		if err != nil {
			http.Error(w, "Failed to store message", http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusCreated, msg)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
