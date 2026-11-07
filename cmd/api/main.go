package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bhavesh256/go-api-project/internal/config"
)

func main() {

	cfg := config.MustLoad()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "all ok"}`))
	})

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

//default text/plain
// default router - globally mux, problem conflicts,
// default time out zero no timeout
// significane of internal package
