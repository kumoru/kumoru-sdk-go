package kumoru

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

func (k *KumoruClient) NewRequest() (*http.Request, error) {
	switch k.Method {
	case POST, PUT, PATCH:
		if k.TargetType == "json" {

			var contentJson []byte

			if len(k.Data) != 0 {
				contentJson, _ = json.Marshal(k.Data)
			}

			contentReader := bytes.NewReader(contentJson)
			req, err := http.NewRequest(k.Method, k.Url, contentReader)
			req.Header.Set("Content-Type", "application/json")
			return req, err
		} else if k.TargetType == "form" {
			req, err := http.NewRequest(k.Method, k.Url, bytes.NewBufferString(k.RawString))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			return req, err
		} else if k.TargetType == "text" {
			req, err := http.NewRequest(k.Method, k.Url, strings.NewReader(k.RawString))
			req.Header.Set("Content-Type", "text/plain")
			return req, err
		} else if k.TargetType == "xml" {
			req, err := http.NewRequest(k.Method, k.Url, strings.NewReader(k.RawString))
			req.Header.Set("Content-Type", "application/xml")
			return req, err
		}
	case GET, HEAD, DELETE:
		return http.NewRequest(k.Method, k.Url, nil)
	}

	return nil, nil
}
