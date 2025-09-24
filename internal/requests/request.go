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

func (r *Request) SetMethod(m string) {
	r.Method = m
}

func (r *Request) SetURL(u string) {
	r.URL = u
}

func (r *Request) SetHeader(h http.Header) {
	r.Header = h
}
