/*
Copyright 2016 Kumoru.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
