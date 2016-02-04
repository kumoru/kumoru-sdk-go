package secrets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

type Secret struct {
	CreatedAt string `"json: created_at"`
	Value     string `"json: value"`
	UpdatedAt string `"json: updated_at"`
	Uuid      string `"json: uuid"`
}

// Create is a Secret method that will create a secret with the specified value
func (s *Secret) Create() (*Secret, *http.Response, []error) {
	k := kumoru.New()

	k.Post(fmt.Sprintf("%v/v1/secrets/", k.EndPoint.Authorization))
	k.Send(genParameters(s.Value))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if errs != nil {
		return s, resp, errs
	}

	err := json.Unmarshal([]byte(body), &s)

	if err != nil {
		errs = append(errs, err)
		return s, resp, errs
	}

	return s, resp, nil
}

// Show is a Secret method will call the appropriate URI and return a specific secret
func (s *Secret) Show(secretUuid *string) (*Secret, *http.Response, []error) {
	secret := Secret{}
	k := kumoru.New()

	k.Get(fmt.Sprintf("%s/v1/secrets/%s", k.EndPoint.Authorization, *secretUuid))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if errs != nil {
		return &secret, resp, errs
	}

	err := json.Unmarshal([]byte(body), &secret)

	if err != nil {
		errs = append(errs, err)
		return &secret, resp, errs
	}

	return &secret, resp, errs
}

func genParameters(value string) string {
	var params string

	if value != "" {
		params += fmt.Sprintf("value=%s&", url.QueryEscape(value))
	}

	return params
}
