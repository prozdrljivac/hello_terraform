// @title           Hello Terraform API
// @version         1.0
// @description     This is a simple API to post and list messages
// @host            localhost:8080
// @BasePath        /
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prozdrljivac/hello_terraform/internal/config"
	"github.com/prozdrljivac/hello_terraform/internal/db/filestorage"
	"github.com/prozdrljivac/hello_terraform/internal/handler"

	_ "github.com/prozdrljivac/hello_terraform/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	file, err := filestorage.NewFileStorage(ctx, "dummy_db.json")
	if err != nil {
		log.Fatalf("failed to connect to file storage: %v", err)
	}
	defer file.Close()

	repo := filestorage.NewFileStorageMessageRepository(file)
	h := handler.NewMessageHandler(repo)

	mux := http.NewServeMux()
	mux.Handle("/", withCORS(h))
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	addr := ":" + cfg.ServerPort
	srv := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server running at http://localhost%s\n", addr)
	fmt.Println("Swagger docs at http://localhost" + addr + "/swagger/index.html")

	log.Fatal(srv.ListenAndServe())
}

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}
