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

	k.Post(fmt.Sprintf("%v/v1/pools/", k.EndPoint.Pool))
	k.Send(fmt.Sprintf("location=%s&credentials=%s", url.QueryEscape(location), url.QueryEscape(credentials)))
	k.SignRequest(true)
	return k.End()
}

func List() (*http.Response, string, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%v/v1/pools/", k.EndPoint.Pool))
	k.SignRequest(true)
	return k.End()
}

func Show(uuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%v/v1/pools/%s", k.EndPoint.Pool, uuid))
	k.SignRequest(true)
	return k.End()
}

func Delete(uuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Delete(fmt.Sprintf("%v/v1/pools/%s", k.EndPoint.Pool, uuid))
	k.SignRequest(true)
	return k.End()
}
