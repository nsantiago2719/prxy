// Package handler contains the http.Handler for the routes
package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/nsantiago2719/prxy/internal/plugins"
	"github.com/nsantiago2719/prxy/internal/requests"
)

// RootHandler returns a simple hello world message
// Required headers are x-prxy-url and x-prxy-method
// Optional custom headers should have a prefix of x-prxy-
// this  optional custom headers will be added to the request
func RootHandler(w http.ResponseWriter, r *http.Request) error {
	// Check if the required headers are present
	if r.Header.Get("x-prxy-url") == "" {
		return errors.New("x-prxy-url header is required")
	}

	// Set initial request values
	request := requests.Init(r.Method, r.Header.Get("x-prxy-url"))

	err := plugins.Init(&request, r.Header)

	resp, err := request.Send()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	// Copy the response body to the client
	io.Copy(w, resp.Body)
	// Close the response body
	defer resp.Body.Close()

	return nil
}

// HealthResponse is a health check response struct
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HealthHandler returns a simple health check message
func HealthHandler(w http.ResponseWriter, _ *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	healthResponse := HealthResponse{
		Status:  "OK",
		Message: "Health check",
	}
	jsonResponse, err := json.Marshal(healthResponse)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(jsonResponse))
	if err != nil {
		return err
	}
	return nil
}
