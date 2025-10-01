// Package routes contains the http.Handler with all the routes
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/nsantiago2719/prxy/internal/handlers"
	"github.com/nsantiago2719/prxy/internal/loggers"
)

type handlerFunc func(http.ResponseWriter, *http.Request) error

// Error is a struct for general error responses
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func makeHandler(f handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log the request for Info
		logger := loggers.NewLogger()
		logger.Info("Request", "method", r.Method, "url", r.URL.Path, "x-prxy-url", r.Header.Get("x-prxy-url"))

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
			logger.Error(err.Message)
		}
	}
}

// NewRoutes returns a new http.Handler with all the routes
func NewRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", makeHandler(handler.RootHandler))
	mux.HandleFunc("/health", makeHandler(handler.HealthHandler))

	return mux
}
