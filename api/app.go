package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	serverPort := getEnvOrSetDefault("SERVER_PORT", "8080")

	messageHandler := &MessageHandler{}
	s := &http.Server{
		Addr:           ":" + serverPort,
		Handler:        messageHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Server is running at http://localhost:%s\n", serverPort)
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

func getEnvOrSetDefault(key string, defaultValue string) string {
	envValue := os.Getenv(key)

	if envValue == "" {
		envValue = defaultValue
	}

	return envValue
}
