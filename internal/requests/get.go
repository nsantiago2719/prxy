package requests

import (
	"net/http"
)

// Get sends a GET request to the backend service
// with the given url and headers
func Get(url string, header http.Header) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
