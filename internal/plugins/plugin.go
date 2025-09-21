package plugins

import (
	"fmt"
	"net/http"
)

func Init(Header http.Header) {
	setCustomHeader(Header)
}
