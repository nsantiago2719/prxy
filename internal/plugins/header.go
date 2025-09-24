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
	if hasValue(header.Get("x-prxy-request-id")) {
		customHeaders.Set("x-prxy-request-id", header.Get("x-prxy-request-id"))
	}
	request.SetHeader(customHeaders)

	return nil
}

func hasValue(s string) bool {
	return s != ""
}
