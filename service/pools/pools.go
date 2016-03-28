package pools

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

// Create a new pool
func Create(location string, credentials string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Post(fmt.Sprintf("%v/v1/pools/", k.EndPoint.Pool))
	k.Send(fmt.Sprintf("location=%s&credentials=%s", url.QueryEscape(location), url.QueryEscape(credentials)))
	k.SignRequest(true)

	return k.End()
}

// Delete a pool that matches the given UUID
func Delete(UUID string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Delete(fmt.Sprintf("%v/v1/pools/%s", k.EndPoint.Pool, UUID))
	k.SignRequest(true)

	return k.End()
}

// List all available pools
func List() (*http.Response, string, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%v/v1/pools/", k.EndPoint.Pool))
	k.SignRequest(true)

	return k.End()
}

// Patch a pool that matches an UUID
func Patch(UUID, credentials string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Patch(fmt.Sprintf("%v/v1/pools/%s", k.EndPoint.Pool, UUID))
	k.Send(fmt.Sprintf("credentials=%s", url.QueryEscape(credentials)))

	k.SignRequest(true)
	return k.End()
}

// Get a pool model that matches the given UUID
func Show(UUID string, wrappedRequest *http.Request) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%v/v1/pools/%s", k.EndPoint.Pool, UUID))
	if wrappedRequest != nil {
		k.ProxyRequest(wrappedRequest)
		if wrappedRequest.Header.Get("X-Kumoru-Context") != "" {
			k.SetHeader("X-Kumoru-Context", wrappedRequest.Header.Get("X-Kumoru-Context"))
		}
	}
	k.SignRequest(true)

	return k.End()
}
