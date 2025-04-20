package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(dsn)

	dbpool, err := pgxpool.New(ctx, dsn)
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
