package kumoru

import (
	"bytes"
	"net/http"
	"strings"
)

// NewRequest sets the appropriate header and appropriate request content
func (k *Client) NewRequest() (*http.Request, error) {
	switch k.Method {
	case POST, PUT, PATCH:
		if k.TargetType == "json" {
			req, err := http.NewRequest(k.Method, k.URL, strings.NewReader(k.RawString))
			req.Header.Set("Content-Type", "application/json")
			return req, err
		} else if k.TargetType == "form" {
			req, err := http.NewRequest(k.Method, k.URL, bytes.NewBufferString(k.RawString))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			return req, err
		} else if k.TargetType == "text" {
			req, err := http.NewRequest(k.Method, k.URL, strings.NewReader(k.RawString))
			req.Header.Set("Content-Type", "text/plain")
			return req, err
		} else if k.TargetType == "xml" {
			req, err := http.NewRequest(k.Method, k.URL, strings.NewReader(k.RawString))
			req.Header.Set("Content-Type", "application/xml")
			return req, err
		}
	case GET, HEAD, DELETE:
		return http.NewRequest(k.Method, k.URL, nil)
	}

	return nil, nil
}
