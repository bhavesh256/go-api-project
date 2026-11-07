package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bhavesh256/go-api-project/internal/config"
)

func main() {
	// load and validate application configuration.
	// note on the 'internal' package: go compiler rules restrict code inside an
	// 'internal' directory so it can only be imported by packages within the same
	// parent module, keeping private implementation details hidden from external users.
	cfg := config.MustLoad()

	// initialize a new request multiplexer (router).
	// unlike the global router (http.DefaultServeMux), creating a custom mux avoids
	// global state pollution and potential security conflicts from imported packages.
	mux := http.NewServeMux()

	// register a route handler using go 1.22+ enhanced routing patterns.
	// it natively supports http methods ("get /healthz"), eliminating the need
	// for manual r.method checks inside the handler.
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		// explicitly set content-type header before writing headers and response body.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "all ok"}`))
	})

	// configure http server timeouts.
	// the standard library's http.server defaults to zero timeouts, which leaves
	// the server vulnerable to resource exhaustion and slow-client attacks (like slowloris).
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10, // max time to read the entire request, including body
		WriteTimeout: time.Second * 30, // max time to write the response
		IdleTimeout:  time.Second * 60, // max time to wait for the next request when keep-alive is enabled
	}

	log.Printf("server running on port:%v", cfg.Port)

	// start the server and block. if listenandserve fails, log the fatal error and exit.
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
