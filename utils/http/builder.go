package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestConfigBuilder struct {
	config *RequestConfig
}

func (rc *RequestConfig) Builder() *RequestConfigBuilder {
	return &RequestConfigBuilder{
		config: &RequestConfig{
			Headers: make(map[string]string),
			Params:  make(map[string]string),
		},
	}
}

func (rb *RequestConfigBuilder) Method(method Method) *RequestConfigBuilder {
	rb.config.Method = method
	return rb
}

func (rb *RequestConfigBuilder) Url(url string) *RequestConfigBuilder {
	rb.config.Url = url
	return rb
}

func (rb *RequestConfigBuilder) AddHeader(key, value string) *RequestConfigBuilder {
	rb.config.Headers[key] = value
	return rb
}

func (rb *RequestConfigBuilder) AddParam(key, value string) *RequestConfigBuilder {
	rb.config.Params[key] = value
	return rb
}

func (rb *RequestConfigBuilder) Body(body []byte) *RequestConfigBuilder {
	rb.config.Body = body
	return rb
}

func (rb *RequestConfigBuilder) Build() *RequestConfig {
	return rb.config
}

func (rc *RequestConfig) NewRequest() (*http.Request, error) {
	switch rc.Method {
	case GET, PUT, POST, DELETE:
		break
	default:
		rc.Method = GET
	}

	if len(rc.Params) > 0 {
		var stringBuilder strings.Builder
		stringBuilder.WriteString(rc.Url)
		stringBuilder.WriteString("?")
		for k, v := range rc.Params {
			stringBuilder.WriteString(k)
			stringBuilder.WriteString("=")
			stringBuilder.WriteString(v)
			stringBuilder.WriteString("&")
		}
		result := stringBuilder.String()
		result = result[:len(result)-1]
		if res, err := url.Parse(result); err == nil {
			rc.Url = res.String()
		}
	}

	request, err := http.NewRequest(string(rc.Method), rc.Url, io.NopCloser(bytes.NewReader(rc.Body)))
	if err != nil {
		return nil, errors.New("fail to create request: " + err.Error())
	}
	if len(rc.Headers) > 0 {
		for k, v := range rc.Headers {
			request.Header.Set(k, v)
		}
	}

	return request, nil

}
