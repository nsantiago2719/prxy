package main

import (
	"net/http"

	"github.com/nsantiago2719/prxy/internal/routes"
)

func main() {
	router := routes.NewRoutes()
	http.ListenAndServe(":8080", router)
}
