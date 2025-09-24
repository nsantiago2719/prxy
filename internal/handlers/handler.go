// Package handler contains the http.Handler for the routes
package handler

import (
	"errors"
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

	// Check if the required headers are present
	if r.Header.Get("x-prxy-url") == "" {
		slog.Error("x-prxy-url header is required")
		return errors.New("x-prxy-url header is required")
	}

	if r.Header.Get("x-prxy-method") == "" {
		slog.Error("x-prxy-method header is required")
		return errors.New("x-prxy-method header is required")
	}

	// Set the method and url from the headers
	// this will be used on sending the request to the backend service
	request.SetMethod(r.Header.Get("x-prxy-method"))
	request.SetURL(r.Header.Get("x-prxy-url"))

	err := plugins.Init(&request, r.Header)

	resp, err := request.Send()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	fmt.Println(resp)
	return nil
}
