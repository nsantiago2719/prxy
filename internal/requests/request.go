package requests

import (
	"net/http"
)

type Request struct {
	Method string
	URL    string
	Header http.Header
}

func Init() Request {
	return Request{
		Method: "",
		URL:    "",
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
	default:
		return nil, nil
	}
}

func (r *Request) SetMethod(m string) {
	r.Method = m
}

func (r *Request) SetURL(u string) {
	r.URL = u
}

func (r *Request) SetHeader(h http.Header) {
	r.Header = h
}
