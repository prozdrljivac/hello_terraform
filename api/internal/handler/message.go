package handler

import (
	"encoding/json"
	"net/http"

	"api/hello_terraform/internal/db"
	"api/hello_terraform/internal/model"
)

var _ = model.Message{}

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
		h.ListMessages(w, r.WithContext(ctx))
	case http.MethodPost:
		h.CreateMessage(w, r.WithContext(ctx))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// @Summary      List messages
// @Description  Retrieve all stored messages
// @Tags         messages
// @Produce      json
// @Success      200 {array} model.Message
// @Failure      500 {string} string "Internal error"
// @Router       / [get]
func (h *MessageHandler) ListMessages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	messages, err := h.repo.ListMessages(ctx)
	if err != nil {
		http.Error(w, "Failed to list messages", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, messages)
}

type createMessageRequest struct {
	Text string `json:"text"`
}

// @Summary      Create a message
// @Description  Store a new message with text
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        message body createMessageRequest true "Message payload"
// @Success      201 {object} model.Message
// @Failure      400 {string} string "Missing or invalid 'text' field"
// @Failure      500 {string} string "Internal error"
// @Router       / [post]
func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req createMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		http.Error(w, "Missing 'text' field", http.StatusBadRequest)
		return
	}

	msg, err := h.repo.CreateMessage(ctx, req.Text)
	if err != nil {
		http.Error(w, "Failed to store message", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusCreated, msg)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
