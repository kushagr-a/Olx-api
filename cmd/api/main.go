package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kushagra/olx-api/internal/config"
	"github.com/kushagra/olx-api/internal/db"
	"github.com/kushagra/olx-api/internal/handlers"
)

/*
go run :- build + executable for devlopment
go build -o bin/main ./cmd/api/main.go :- build for production bin folder main file and build main.go
./bin/main :- this is command for execute build file.
Makefile :- used as a shortcut
Makefile :- running like command make build, make run, make clean

*/

// go main server
func main() {

	// config package
	cfg := config.MustLoad()

	// db connection
	_, err := db.ConnectDb(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("db connection failed: %v\n", err)
	}

	log.Println("db connected successfully")

	// custom router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.Healthz)

	// server creation and configuration
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second, // 10 sec min time to read request
		WriteTimeout: 30 * time.Second, // 10 sec min time to write response
		IdleTimeout:  60 * time.Second, // 10 sec min time to id
	}
	// print log for server start and port
	log.Println("API is running on http://localhost:" + cfg.Port)

	// Listening
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v\n", err) // fatalf means formated log
	}
}
