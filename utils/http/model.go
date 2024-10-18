package http

import "net/http"

type Method string

const (
	POST   Method = "POST"
	GET    Method = "GET"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type RequestConfig struct {
	Method  Method
	Url     string
	Headers map[string]string
	Params  map[string]string
	Body    []byte
}

type RequestInterceptor interface {
	InterceptRequest(*http.Request) error
}

type ResponseInterceptor interface {
	InterceptResponse(*http.Response) error
}

type Handler[T any] func(*http.Response) (T, error)

type Client[T any] struct {
	Instance             *http.Client
	RequestInterceptors  []RequestInterceptor
	ResponseInterceptors []ResponseInterceptor
}
