package plugins

import (
	"net/http"

	"github.com/nsantiago2719/prxy/internal/requests"
)

func Init(request *requests.Request, header http.Header) error {
	err := setHeader(request, header)

	return err
}
