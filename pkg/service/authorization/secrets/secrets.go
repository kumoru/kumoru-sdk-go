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

package secrets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kumoru/kumoru-sdk-go/pkg/kumoru"
)

type Secret struct {
	CreatedAt string   `"json: created_at"`
	Labels    []string `"json: labels,omitempty"`
	UpdatedAt string   `"json: updated_at"`
	Uuid      string   `"json: uuid"`
	Value     string   `"json: value"`
}

// Create is a Secret method that will create a secret with the specified value
func (s *Secret) Create() (*Secret, *http.Response, []error) {
	k := kumoru.New()

	k.Post(fmt.Sprintf("%v/v1/secrets/", k.EndPoint.Authorization))
	k.Send(genParameters(s.Value, s.Labels))
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

//List retreives all secrets a role has access to
func List() ([]*Secret, *http.Response, []error) {
	apps := []*Secret{}
	k := kumoru.New()

	k.Get(fmt.Sprintf("%s/v1/secrets/", k.EndPoint.Authorization))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if len(errs) > 0 {
		return nil, resp, errs
	}

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
	}

	err := json.Unmarshal([]byte(body), &apps)

	if err != nil {
		errs = append(errs, fmt.Errorf("%s", err))
	}

	return apps, resp, nil
}

//Helpers
func genParameters(value string, labels []string) string {
	var params string

	if value != "" {
		params += fmt.Sprintf("value=%s&", url.QueryEscape(value))
	}

	for _, v := range labels {
		params += fmt.Sprintf("labels=%s&", url.QueryEscape(v))
	}

	return params
}
