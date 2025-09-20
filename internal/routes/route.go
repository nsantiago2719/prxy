// Package routes contains the http.Handler with all the routes
package routes

import (
	"net/http"

	"github.com/nsantiago2719/prxy/internal/handlers"
)

// NewRoutes returns a new http.Handler with all the routes
func NewRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.RootHandler)

	return mux
}
