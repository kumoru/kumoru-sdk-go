/*
Copyright 2016 Kumoru.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
See the License for the specific language governing permissions and
limitations under the License.
*/

package deployments

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kumoru/kumoru-sdk-go/pkg/kumoru"
)

type Deployment struct {
	ApplicationUUID string                 `json:"application_uuid"`
	CreatedAt       string                 `json:"created_at"`
	Environment     map[string]string      `json:"environment"`
	ImageId         string                 `json:"image_id"`
	ImageTag        string                 `json:"tag"`
	ImageUrl        string                 `json:"image_url"`
	Metadata        map[string]interface{} `json:"metadata"`
	Ports           []string               `json:"ports"`
	SSLPorts        []string               `json:"ssl_ports"`
	Url             string                 `json:"url"`
	Uuid            string                 `json:"uuid"`
}

// List is a method will call the appropriate URI and return a list of all deployments
func (d *Deployment) List(applicationUuid string) (*[]Deployment, *http.Response, []error) {
	deployments := []Deployment{}
	k := kumoru.New()

	k.Get(fmt.Sprintf("%s/v1/applications/%s/deployments/", k.EndPoint.Application, applicationUuid))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if errs != nil {
		return &deployments, resp, errs
	}

	err := json.Unmarshal([]byte(body), &deployments)

	if err != nil {
		errs = append(errs, err)
		return &deployments, resp, errs
	}

	return &deployments, resp, nil
}

// Show is a method will call the appropriate URI and return a specific deployment
func (d *Deployment) Show(applicationUuid, deploymentUuid string) (*Deployment, *http.Response, []error) {
	deployment := Deployment{}
	k := kumoru.New()

	k.Get(fmt.Sprintf("%s/v1/applications/%s/deployments/%s", k.EndPoint.Application, applicationUuid, deploymentUuid))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if errs != nil {
		return &deployment, resp, errs
	}

	err := json.Unmarshal([]byte(body), &deployment)

	if err != nil {
		errs = append(errs, err)
		return &deployment, resp, errs
	}

	return &deployment, resp, errs
}
