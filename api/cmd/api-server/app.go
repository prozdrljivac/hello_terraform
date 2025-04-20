package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"api/hello_terraform/internal/config"
	"api/hello_terraform/internal/db"
	"api/hello_terraform/internal/handler"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbpool, err := db.NewPostgresPool(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer dbpool.Close()
	_, err = dbpool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS messages (
			id SERIAL PRIMARY KEY,
			text TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create messages table: %v", err)
	}

	repo := db.NewPostgresMessageRepository(dbpool)

	handler := handler.NewMessageHandler(repo)

	addr := ":" + cfg.ServerPort
	srv := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server running at http://localhost%s\n", addr)
	log.Fatal(srv.ListenAndServe())
}
