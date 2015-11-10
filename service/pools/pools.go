package pools

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

func Create(location string, credentials string) (*http.Response, string, []error) {
	k := kumoru.New()

	if k.Tokens.Public == "" || k.Tokens.Private == "" {
		k.Logger.Fatal("Update your config with a set of tokens")
	}

	return k.Post(fmt.Sprintf("%v/v1/pools/", k.EndPoint.Pool)).
		Send(fmt.Sprintf("location=%s&credentials=%s", url.QueryEscape(location), url.QueryEscape(credentials))).
		SignRequest(true).
		End()
}

func List() (*http.Response, string, []error) {
	k := kumoru.New()

	return k.Get(fmt.Sprintf("%v/v1/pools/", k.EndPoint.Pool)).
		SignRequest(true).
		End()
}
