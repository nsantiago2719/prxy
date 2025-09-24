// Package handler contains the http.Handler for the routes
package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/nsantiago2719/prxy/internal/plugins"
	"github.com/nsantiago2719/prxy/internal/requests"
)

// RootHandler returns a simple hello world message
// Required headers are x-prxy-url and x-prxy-method
// Optional custom headers should have a prefix of x-prxy-
// this  optional custom headers will be added to the request
func RootHandler(w http.ResponseWriter, r *http.Request) error {
	request := requests.Init()

	request.SetMethod(r.Header.Get("x-prxy-method"))
	request.SetURL(r.Header.Get("x-prxy-url"))

	err := plugins.Init(&request, r.Header)
	if err != nil {
		slog.Error(err.Error())
	}
	return nil
}
