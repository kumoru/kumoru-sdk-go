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

package pools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kumoru/kumoru-sdk-go/pkg/kumoru"
)

type Location struct {
	AggregateResources map[string]float32 `json:"aggregate_resources"`
	CreatedAt          string             `json:"created_at"`
	Identifier         string             `json:"location"`
	Provider           string             `json:"provider"`
	PoolId             string             `json:"stack_id"`
	Status             string             `json:"status"`
	UpdatedAt          string             `json:"updated_at"`
	Url                string             `json:"url"`
	Uuid               string             `json:"uuid"`
	ApiVersion         string             `json:"api_version"`
}

// Create is a method on a Location that will create Kumoru resources in the provider region
func (l *Location) Create() (*Location, *http.Response, []error) {
	k := kumoru.New()

	k.Post(fmt.Sprintf("%v/v1/pools/", k.EndPoint.Pool))
	k.Send(fmt.Sprintf("location=%s", url.QueryEscape(l.Identifier)))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
		return l, resp, errs
	}

	err := json.Unmarshal([]byte(body), &l)

	if err != nil {
		errs = append(errs, err)
		return l, resp, errs
	}

	return l, resp, nil
}

//Delete is a method on a Location that will remove Kumoru resources from the provider region
func (l *Location) Delete(uuid string) (*Location, *http.Response, []error) {
	k := kumoru.New()

	k.Delete(fmt.Sprintf("%v/v1/pools/%s", k.EndPoint.Pool, uuid))
	k.SignRequest(true)

	resp, _, errs := k.End()

	if errs != nil {
		return l, resp, errs
	}

	return l, resp, nil
}

//List is a method on a Location that will list all Locations an user has access to
func (l *Location) List() (*[]Location, *http.Response, []error) {
	locations := []Location{}
	k := kumoru.New()

	k.Get(fmt.Sprintf("%v/v1/pools/", k.EndPoint.Pool))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
		return &locations, resp, errs
	}

	err := json.Unmarshal([]byte(body), &locations)

	if err != nil {
		errs = append(errs, err)
		return &locations, resp, errs
	}

	return &locations, resp, nil
}

//Show is a method on a Location that will show all the details of a particular Location
func (l *Location) Show() (*Location, *http.Response, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%v/v1/pools/%s", k.EndPoint.Pool, l.Uuid))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
		return l, resp, errs
	}

	err := json.Unmarshal([]byte(body), l)

	if err != nil {
		errs = append(errs, err)
		return l, resp, errs
	}

	return l, resp, errs

}
