package main

import (
	v1 "finance-tracker-backend/internal/api/v1"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	v1API := v1.NewRoute()

	v1API.RegisterHandlers(mux)

	srv := http.Server{
		Addr:         ":3000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
