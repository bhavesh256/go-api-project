package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bhavesh256/go-api-project/internal/config"
	"github.com/bhavesh256/go-api-project/internal/db"
	"github.com/bhavesh256/go-api-project/internal/handlers"
)

func main() {
	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}
	log.Println("database connected")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /listings", handlers.List(db))
	mux.HandleFunc("DELETE /listings/{id}", handlers.DeleteListing(db))

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("server running on port:%v", cfg.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
