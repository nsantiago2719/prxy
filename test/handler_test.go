package main_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nsantiago2719/prxy/internal/handlers"
	"github.com/nsantiago2719/prxy/internal/routes"
)

func TestRootHandler(t *testing.T) {
	// TODO: add more tests
	tests := []struct {
		name       string
		method     string
		prxyID     string
		status     int
		wantPrxyID string
		wantErr    bool
	}{
		{
			name:       "TestRootHandler_GetSuccess",
			method:     http.MethodGet,
			prxyID:     "123123123",
			status:     http.StatusOK,
			wantPrxyID: "123123123",
			wantErr:    false,
		},
		{
			name:       "TestRootHandler_EmptyPrxyID",
			method:     http.MethodGet,
			prxyID:     "",
			status:     http.StatusOK,
			wantPrxyID: "",
			wantErr:    true,
		},
		{
			name:       "TestRootHandler_PostSuccess",
			method:     http.MethodPost,
			prxyID:     "123123123",
			status:     http.StatusOK,
			wantPrxyID: "123123123",
			wantErr:    false,
		},
	}

	// backend test server
	mockServiceHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prxyID := r.Header.Get("x-prxy-request-id")
		response := Response{
			PrxyID: prxyID,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))
		w.Header().Set("x-prxy-request-id", prxyID)
	})
	testServer := httptest.NewServer(mockServiceHandler)
	defer testServer.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(makeHandler(handler.RootHandler))

			backendTestServer := testServer.URL
			backendXProxyID := tt.prxyID

			// Set the backend url and request id in the request header
			req.Header.Set("x-prxy-url", backendTestServer)
			req.Header.Set("x-prxy-request-id", backendXProxyID)

			handler.ServeHTTP(rr, req)

			body, err := io.ReadAll(rr.Body)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			var response Response
			err = json.Unmarshal(body, &response)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			if !tt.wantErr {
				if response.PrxyID != tt.wantPrxyID {
					t.Errorf("Error: prxy-id is not equal to %s", tt.wantPrxyID)
				}
			}

			if tt.wantErr {
				if response.PrxyID != tt.wantPrxyID {
					t.Errorf("Error: prxy-id is not empty")
				}
			}
		})
	}
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func TestHealthHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		wantStatus string
	}{
		{
			name:       "TestHealthHandler_Success",
			method:     http.MethodGet,
			wantStatus: "OK",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/health", nil)

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(makeHandler(handler.HealthHandler))

			handler.ServeHTTP(rr, req)

			body, err := io.ReadAll(rr.Body)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			handler.ServeHTTP(rr, req)
			body, err = io.ReadAll(rr.Body)

			var response HealthResponse
			err = json.Unmarshal(body, &response)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			if response.Status != tt.wantStatus {
				t.Errorf("Error: prxy-id is not equal to %s", tt.wantStatus)
			}
		})
	}
}

// Response is a simple response struct for testing
type Response struct {
	PrxyID string `json:"prxy-id"`
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

func makeHandler(f handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the handler function
		if err := f(w, r); err != nil {
			// Log the error and return the error
			err := routes.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
		}
	}
}
