package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (client *Client[T]) Do(request *http.Request, handler Handler[T], config *RequestConfig) (*T, error) {

	if config != nil {
		if len(config.Method) > 0 {
			request.Method = string(config.Method)
		}

		if len(config.Url) > 0 {
			if res, err := url.Parse(config.Url); err != nil {
				return nil, errors.New("invalid url")
			} else {
				request.URL = res
			}
		}

		if len(config.Body) > 0 {
			request.Body = io.NopCloser(bytes.NewReader(config.Body))
		}

		if len(config.Headers) > 0 {
			for key, value := range config.Headers {
				request.Header.Set(key, value)
			}
		}

		if len(config.Params) > 0 {
			var stringBuilder strings.Builder
			stringBuilder.WriteString(request.URL.String())
			stringBuilder.WriteString("?")
			for key, value := range config.Params {
				stringBuilder.WriteString(key)
				stringBuilder.WriteString("=")
				stringBuilder.WriteString(value)
				stringBuilder.WriteString("&")
			}
			result := stringBuilder.String()
			result = result[:len(result)-1]
			if res, err := url.Parse(result); err == nil {
				request.URL = res
			}
		}

	}

	for _, requestInterceptor := range client.RequestInterceptors {
		err := requestInterceptor.InterceptRequest(request)
		if err != nil {
			return nil, errors.New("fatal error while intercept request")
		}
	}
	response, err := client.Instance.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if t, err := handler(response); err != nil {
		return nil, err
	} else {
		return &t, nil
	}
}
