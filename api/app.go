package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	messageHandler := &MessageHandler{}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        messageHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(s.ListenAndServe())
}

type MessageHandler struct{}

func (h *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Returning list of messages (TODO)")
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Storing message (TODO)")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
