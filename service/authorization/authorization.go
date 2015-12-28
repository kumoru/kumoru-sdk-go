package authorization

import (
	"fmt"
	"net/http"

	"code.google.com/p/go-uuid/uuid"
	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

func GetTokens(username, password string) (string, *http.Response, string, []error) {
	k := kumoru.New()

	token := uuid.New()

	k.Put(fmt.Sprintf("%v/v1/tokens/%v", k.EndPoint.Authorization, token))
	k.SetBasicAuth(username, password)
	resp, body, errs := k.End()

	return token, resp, body, errs
}

func CreateAcct(username, given_name, surname, password string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Put(fmt.Sprintf("%s/v1/accounts/%s", k.EndPoint.Authorization, username))
	k.Send(fmt.Sprintf("given_name=%s&surname=%s&password=%s", given_name, surname, password))
	return k.End()
}

func ShowAcct(username string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%v/v1/accounts/%v", k.EndPoint.Authorization, username))
	return k.End()
}
