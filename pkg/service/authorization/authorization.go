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

//Package authorization provides an Account type and pertinent methods for this type.
package authorization

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kumoru/kumoru-sdk-go/pkg/kumoru"
	"github.com/pborman/uuid"
)

//Account represents a user and pretinent metadata about that user.
type Account struct {
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	GivenName string `json:"given_name"`
	RoleUUID  string `json:"role_uuid"`
	Surname   string `json:"surname"`
	UpdatedAt string `json:"updated_at"`
}

//CreateAcct requests a particular account be made in Kumoru.
//It returns the updated Account.
func (a *Account) CreateAcct(password string) (*Account, *http.Response, []error) {
	k := kumoru.New()

	k.Put(fmt.Sprintf("%s/v1/accounts/%s", k.EndPoint.Authorization, a.Email))
	k.Send(fmt.Sprintf("given_name=%s&surname=%s&password=%s", a.GivenName, a.Surname, password))

	resp, body, errs := k.End()

	if len(errs) > 0 {
		return a, resp, errs
	}

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
	}

	err := json.Unmarshal([]byte(body), &a)

	if err != nil {
		errs = append(errs, err)
		return a, resp, errs
	}

	return a, resp, nil
}

//ResetPassword requests the password be reset for a given Account.
func (a *Account) ResetPassword() (*Account, *http.Response, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%v/v1/accounts/%v/password/resets/", k.EndPoint.Authorization, a.Email))
	resp, _, errs := k.End()

	return a, resp, errs
}

//Show requests account details from Kumoru and marshals the data into the Account type.
func (a *Account) Show() (*Account, *http.Response, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%v/v1/accounts/%v", k.EndPoint.Authorization, a.Email))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if len(errs) > 0 {
		return a, resp, errs
	}

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
		return a, resp, errs
	}

	err := json.Unmarshal([]byte(body), &a)

	if err != nil {
		errs = append(errs, err)
		return a, resp, errs
	}

	return a, resp, nil
}

//GetTokens generates a new token(uuid), stores this token in Kumoru and retrieves the private half of the token.
func GetTokens(username, password string) (string, *http.Response, string, []error) {
	k := kumoru.New()

	token := uuid.New()

	k.Put(fmt.Sprintf("%v/v1/tokens/%v", k.EndPoint.Authorization, token))
	k.SetBasicAuth(username, password)
	resp, body, errs := k.End()

	return token, resp, body, errs
}
