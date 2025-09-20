// Package handler contains the http.Handler for the routes
package handler

import (
	"net/http"
)

// RootHandler returns a simple hello world message
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
