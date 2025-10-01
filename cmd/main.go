package main

import (
	"net/http"
	"time"

	"github.com/nsantiago2719/prxy/internal/routes"
)

func main() {
	server := configureServer()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func configureServer() http.Server {
	handlers := routes.NewRoutes()

	return http.Server{
		Addr:    ":8080",
		Handler: handlers,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
