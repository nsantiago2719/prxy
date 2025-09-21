// Package routes contains the http.Handler with all the routes
package routes

import (
	"net/http"

	"github.com/nsantiago2719/prxy/internal/handlers"
)

type handlerFunc func(http.ResponseWriter, *http.Request) error

func makeHandler(f handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the handler function
		if err := f(w, r); err != nil {
			// Log the error
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// NewRoutes returns a new http.Handler with all the routes
func NewRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", makeHandler(handler.RootHandler))

	return mux
}
