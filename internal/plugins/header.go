package plugins

import (
	"net/http"

	"github.com/nsantiago2719/prxy/internal/requests"
)

func setHeader(request *requests.Request, header http.Header) error {
	if len(header) == 0 {
		return nil
	}

	customHeaders := http.Header{}
	// x-prxy-request-id is an optional custom header
	// which is added to the request and will be used by the backend  service
	// to identify the request from frontend
	if hasValue(header.Get("x-prxy-request-id")) {
		customHeaders.Set("x-prxy-request-id", header.Get("x-prxy-request-id"))
	}
	request.SetHeader(customHeaders)

	return nil
}

func hasValue(s string) bool {
	return s != ""
}
