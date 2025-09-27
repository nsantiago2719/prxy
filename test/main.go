package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response is a simple response struct for testing
type Response struct {
	PrxyID string `json:"prxy-id"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
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
	http.ListenAndServe(":3000", nil)
}
