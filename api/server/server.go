package server

import (
	"net/http"
	"time"
)

func RunServer() {
	server := http.Server{
		Addr:         ":12130",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handler(),
	}

	server.ListenAndServe()
}
