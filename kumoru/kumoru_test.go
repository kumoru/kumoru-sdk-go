package kumoru

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// testing for Get method
func TestKumoruGet(t *testing.T) {
	os.Clearenv()
	os.Setenv("KUMORU_CONFIG", "example-cfg.ini")
	const case1_empty = "/"
	const case2_set_header = "/set_header"

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
			if r.Header.Get("API-Key") != "fookey" {
				t.Errorf("Expected 'API-Key' == %q; got %q", "fookey", r.Header.Get("API-Key"))
			}
		}
	}))

	defer ts.Close()

	New().Get(ts.URL + case1_empty).
		End()

	New().Get(ts.URL+case2_set_header).
		SetHeader("API-Key", "fookey").
		End()

	os.Clearenv()
}
