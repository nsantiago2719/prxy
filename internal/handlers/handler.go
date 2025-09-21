// Package handler contains the http.Handler for the routes
package handler

import (
	"net/http"

	"github.com/nsantiago2719/prxy/internal/plugins"
)

type Request struct {
	Method string
	URL    string
	Header http.Header
}

// RootHandler returns a simple hello world message
func RootHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/plain")

	headers := r.Header

	plugins.Init(headers)
	w.Write([]byte("Hello World!"))

	return nil
}
