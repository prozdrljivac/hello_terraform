package main

import (
	"api/hello_terraform/internal/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, cfg.DSN())
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	var msg string
	err = dbpool.QueryRow(ctx, "SELECT 'Hello, world!'").Scan(&msg)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}

	fmt.Println("Message from DB:", msg)

	messageHandler := &MessageHandler{}
	s := &http.Server{
		Addr:           ":" + cfg.ServerPort,
		Handler:        messageHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Server is running at http://localhost:%s\n", cfg.ServerPort)
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
