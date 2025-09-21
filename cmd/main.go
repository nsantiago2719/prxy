package main

import (
	"net/http"
	"time"

	"github.com/nsantiago2719/prxy/internal/routes"
)

func main() {
	router := routes.NewRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
