package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bhavesh256/go-api-project/internal/config"
	"github.com/bhavesh256/go-api-project/internal/handlers"
)

func main() {
	cfg := config.MustLoad()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("GET /home", handlers.Home)

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
