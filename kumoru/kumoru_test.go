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
	const case1_empty = "/v1/pools/"
	const case2_set_header = "/v1/applications/"

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
		case case1_empty:
			t.Logf("case %v ", case1_empty)
		case case2_set_header:
			t.Logf("case %v ", case2_set_header)
			if r.Header.Get("Custom-Header") != "fookey" {
				t.Errorf("Expected 'Custom-Header' == %q; got %q", "fookey", r.Header.Get("Custom-Header"))
			}
		}
	}))

	defer ts.Close()

	New().Get(ts.URL + case1_empty).
		End()

	New().Get(ts.URL+case2_set_header).
		SetHeader("Custom-Header", "fookey").
		End()

	os.Clearenv()
}

// testing for Post method
func TestKumoruPost(t *testing.T) {
	os.Clearenv()
	os.Setenv("KUMORU_CONFIG", "example-cfg.ini")
	const (
		case1_empty      = "/v1/pools/"
		case2_set_header = "/v1/applications/"
		case3_set_query  = "/v1/accounts/"
		case4_send       = "/v1/accounts/victor@kumoru.io"
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
		case case1_empty:
			t.Logf("case %v ", case1_empty)
		case case2_set_header:
			t.Logf("case %v ", case2_set_header)
			if r.Header.Get("Custom-Header") != "fookey" {
				t.Errorf("Expected 'Custom-Header' == %q; got %q", "fookey", r.Header.Get("Custom-Header"))
			}
		case case3_set_query:
			t.Logf("case %v ", case3_set_query)
			v := r.URL.Query()
			if v["rules"][0] != "latest:100" {
				t.Error("Expected latest:100", "| but got", v["rules"][0])
			}
		case case4_send:
			t.Logf("case %v ", case4_send)
			defer r.Body.Close()
			body, _ := ioutil.ReadAll(r.Body)
			if string(body) != "rules=latest:100" {
				t.Error("Expected Body with \"rules=latest:100\"", "| but got", string(body))
			}
		}
	}))

	defer ts.Close()

	New().Post(ts.URL + case1_empty).
		End()

	New().Post(ts.URL+case2_set_header).
		SetHeader("Custom-Header", "fookey").
		End()

	New().Post(ts.URL + case3_set_query).
		Query("rules=latest:100").
		End()

	New().Post(ts.URL + case4_send).
		Send("rules=latest:100").
		End()

	os.Clearenv()
}

// testing for Put method
func TestKumoruPut(t *testing.T) {
	os.Clearenv()
	os.Setenv("KUMORU_CONFIG", "example-cfg.ini")
	const (
		case1_empty      = "/v1/pools/"
		case2_set_header = "/v1/applications/"
		case3_set_query  = "/v1/accounts/"
		case4_send       = "/v1/accounts/victor@kumoru.io"
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
		case case1_empty:
			t.Logf("case %v ", case1_empty)
		case case2_set_header:
			t.Logf("case %v ", case2_set_header)
			if r.Header.Get("Custom-Header") != "fookey" {
				t.Errorf("Expected 'Custom-Header' == %q; got %q", "fookey", r.Header.Get("Custom-Header"))
			}
		case case3_set_query:
			t.Logf("case %v ", case3_set_query)
			v := r.URL.Query()
			if v["rules"][0] != "latest:100" {
				t.Error("Expected latest:100", "| but got", v["rules"][0])
			}
		case case4_send:
			t.Logf("case %v ", case4_send)
			defer r.Body.Close()
			body, _ := ioutil.ReadAll(r.Body)
			if string(body) != "rules=latest:100" {
				t.Error("Expected Body with \"rules=latest:100\"", "| but got", string(body))
			}
		}
	}))

	defer ts.Close()

	New().Put(ts.URL + case1_empty).
		End()

	New().Put(ts.URL+case2_set_header).
		SetHeader("Custom-Header", "fookey").
		End()

	New().Put(ts.URL + case3_set_query).
		Query("rules=latest:100").
		End()

	New().Put(ts.URL + case4_send).
		Send("rules=latest:100").
		End()

	os.Clearenv()
}
