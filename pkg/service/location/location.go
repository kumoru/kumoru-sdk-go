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

package location

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kumoru/kumoru-sdk-go/pkg/kumoru"
)

//Location represents a set of resources in a cloud provider at a given region.
type Location struct {
	OrchestrationURL string `json:"kubernetes_api_url"`
	Provider         string `json:"provider"`
	Region           string `json:"region"`
}

//Create is a method which will request a location be created
func (l *Location) Create() (*Location, *http.Response, []error) {
	k := kumoru.New()

	k.Put(fmt.Sprintf("%s/v1/locations/%s/%s", k.EndPoint.Location, l.Provider, l.Region))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
	}

	err := json.Unmarshal([]byte(body), &l)

	if err != nil {
		errs = append(errs, err)
		return l, resp, errs
	}

	return l, resp, nil
}

//Delete will request that a particular Location be removed
func (l *Location) Delete() (*Location, *http.Response, []error) {
	k := kumoru.New()

	k.Delete(fmt.Sprintf("%s/v1/locations/%s/%s", k.EndPoint.Location, l.Provider, l.Region))
	k.SignRequest(true)

	resp, _, errs := k.End()

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("s", resp.Status))
	}

	return l, resp, nil
}
