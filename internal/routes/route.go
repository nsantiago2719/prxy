// Package routes contains the http.Handler with all the routes
package routes

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/nsantiago2719/prxy/internal/handlers"
)

type handlerFunc func(http.ResponseWriter, *http.Request) error

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func makeHandler(f handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the handler function
		if err := f(w, r); err != nil {
			// Log the error and return the error
			err := Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			slog.Error(err.Message)
		}
	}
}

// NewRoutes returns a new http.Handler with all the routes
func NewRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", makeHandler(handler.RootHandler))

	return mux
}
