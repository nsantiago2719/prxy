package requests

import (
	"errors"
	"net/http"
)

type Request struct {
	Method string
	URL    string
	Header http.Header
}

func Init(method string, url string) Request {
	return Request{
		Method: method,
		URL:    url,
		Header: http.Header{},
	}
}

// Send sends the request to the backend service
// depending on the method set in the request
func (r *Request) Send() (*http.Response, error) {
	switch r.Method {
	case "GET":
		resp, err := Get(r.URL, r.Header)
		if err != nil {
			return nil, err
		}
		return resp, nil
	// TODO: add other methods here
	default:
		err := errors.New("Method is not supported")
		return nil, err
	}
}

func (r *Request) SetHeader(h http.Header) {
	r.Header = h
}
