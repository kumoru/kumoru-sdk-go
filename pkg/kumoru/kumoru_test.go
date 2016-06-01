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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// testing for Get method
func TestKumoruGet(t *testing.T) {
	os.Clearenv()
	os.Setenv("KUMORU_CONFIG", "example-cfg.ini")
	const case1Empty = "/v1/pools/"
	const case2SetHeader = "/v1/applications/"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check method is GET before going to check other features
		if r.Method != GET {
			t.Errorf("Expected method %q; got %q", GET, r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}
		switch r.URL.Path {

		default:
			t.Errorf("No testing for this case yet : %q", r.URL.Path)
		case case1Empty:
			t.Logf("case %v ", case1Empty)
		case case2SetHeader:
			t.Logf("case %v ", case2SetHeader)
			if r.Header.Get("Custom-Header") != "fookey" {
				t.Errorf("Expected 'Custom-Header' == %q; got %q", "fookey", r.Header.Get("Custom-Header"))
			}
		}
	}))

	defer ts.Close()

	k := New()
	k.Get(ts.URL + case1Empty)
	k.End()

	k.Get(ts.URL + case2SetHeader)
	k.SetHeader("Custom-Header", "fookey")
	k.End()

	os.Clearenv()
}

// testing for Post method
func TestKumoruPost(t *testing.T) {
	os.Clearenv()
	os.Setenv("KUMORU_CONFIG", "example-cfg.ini")
	const (
		case1Empty     = "/v1/pools/"
		case2SetHeader = "/v1/applications/"
		case3SetQuery  = "/v1/accounts/"
		case4Send      = "/v1/accounts/victor@kumoru.io"
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check method is POST before going to check other features
		if r.Method != POST {
			t.Errorf("Expected method %q; got %q", POST, r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}
		switch r.URL.Path {

		default:
			t.Errorf("No testing for this case yet : %q", r.URL.Path)
		case case1Empty:
			t.Logf("case %v ", case1Empty)
		case case2SetHeader:
			t.Logf("case %v ", case2SetHeader)
			if r.Header.Get("Custom-Header") != "fookey" {
				t.Errorf("Expected 'Custom-Header' == %q; got %q", "fookey", r.Header.Get("Custom-Header"))
			}
		case case3SetQuery:
			t.Logf("case %v ", case3SetQuery)
			v := r.URL.Query()
			if v["rules"][0] != "latest:100" {
				t.Error("Expected latest:100", "| but got", v["rules"][0])
			}
		case case4Send:
			t.Logf("case %v ", case4Send)
			defer r.Body.Close()
			body, _ := ioutil.ReadAll(r.Body)
			if string(body) != "rules=latest:100" {
				t.Error("Expected Body with \"rules=latest:100\"", "| but got", string(body))
			}
		}
	}))

	defer ts.Close()

	k := New()
	k.Post(ts.URL + case1Empty)
	k.End()

	k.Post(ts.URL + case2SetHeader)
	k.SetHeader("Custom-Header", "fookey")
	k.End()

	k.Post(ts.URL + case3SetQuery)
	k.Query("rules=latest:100")
	k.End()

	k.Post(ts.URL + case4Send)
	k.Send("rules=latest:100")
	k.End()

	os.Clearenv()
}

// testing for Put method
func TestKumoruPut(t *testing.T) {
	os.Clearenv()
	os.Setenv("KUMORU_CONFIG", "example-cfg.ini")
	const (
		case1Empty     = "/v1/pools/"
		case2SetHeader = "/v1/applications/"
		case3SetQuery  = "/v1/accounts/"
		case4Send      = "/v1/accounts/victor@kumoru.io"
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check method is PUT before going to check other features
		if r.Method != PUT {
			t.Errorf("Expected method %q; got %q", PUT, r.Method)
		}
		if r.Header == nil {
			t.Errorf("Expected non-nil request Header")
		}
		switch r.URL.Path {

		default:
			t.Errorf("No testing for this case yet : %q", r.URL.Path)
		case case1Empty:
			t.Logf("case %v ", case1Empty)
		case case2SetHeader:
			t.Logf("case %v ", case2SetHeader)
			if r.Header.Get("Custom-Header") != "fookey" {
				t.Errorf("Expected 'Custom-Header' == %q; got %q", "fookey", r.Header.Get("Custom-Header"))
			}
		case case3SetQuery:
			t.Logf("case %v ", case3SetQuery)
			v := r.URL.Query()
			if v["rules"][0] != "latest:100" {
				t.Error("Expected latest:100", "| but got", v["rules"][0])
			}
		case case4Send:
			t.Logf("case %v ", case4Send)
			defer r.Body.Close()
			body, _ := ioutil.ReadAll(r.Body)
			if string(body) != "rules=latest:100" {
				t.Error("Expected Body with \"rules=latest:100\"", "| but got", string(body))
			}
		}
	}))

	defer ts.Close()

	k := New()
	k.Put(ts.URL + case1Empty)
	k.End()

	k.Put(ts.URL + case2SetHeader)
	k.SetHeader("Custom-Header", "fookey")
	k.End()

	k.Put(ts.URL + case3SetQuery)
	k.Query("rules=latest:100")
	k.End()

	k.Put(ts.URL + case4Send)
	k.Send("rules=latest:100")
	k.End()

	os.Clearenv()
}
