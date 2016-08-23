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

//Package application provides an Application type and pertinent methods for it.
package application

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kumoru/kumoru-sdk-go/pkg/kumoru"
	"github.com/mattbaird/jsonpatch"
)

//Application type represents an Application in Kumoru.
type Application struct {
	Addresses          []string               `json:"addresses,omitempty"`
	CreatedAt          string                 `json:"created_at,omitempty"`
	CurrentDeployments map[string]string      `json:"current_deployments,omitempty"`
	DeploymentToken    string                 `json:"deployment_token,omitempty"`
	Environment        map[string]string      `json:"environment,omitempty"`
	Hash               string                 `json:"hash,omitempty"`
	ImageURL           string                 `json:"image_url"`
	Location           Location               `json:"location"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
	Name               string                 `json:"name"`
	OwnerUUID          string                 `json:"owner_uuid,omitempty"`
	Ports              []string               `json:"ports,omitempty"`
	Rules              map[string]int         `json:"rules,omitempty"`
	SSLPorts           []string               `json:"ssl_ports,omitempty"`
	Status             string                 `json:"status,omitempty"`
	UpdatedAt          string                 `json:"updated_at,omitempty"`
	URL                string                 `json:"url,omitempty"`
	UUID               string                 `json:"uuid,omitempty"`
	APIVersion         string                 `json:"api_version,omitempty"`
	Certificates       Certificates           `json:"certificates,omitempty"`
}

//Location holds pertinent information about the Location the application is deployed in.
type Location struct {
	Provider string `json:"provider,omitempty"`
	Region   string `json:"region,omitempty"`
}

//Certificates type which represents the SSL certificate to be used for an Application.
type Certificates struct {
	Certificate      string `json:"certificate,omitempty"`
	PrivateKey       string `json:"private_key,omitempty"`
	CertificateChain string `json:"certificate_chain,omitempty"`
}

//Application Methods

//Create is a method on an Application which requests that the application be drafted in Kumoru.
func (a *Application) Create() (*Application, *http.Response, []error) {
	var errs []error
	k := kumoru.New()

	k.Post(fmt.Sprintf("%s/v1/applications/", k.EndPoint.Application))
	k.TargetType = "json"
	s, err := json.Marshal(*a)

	if err != nil {
		errs = append(errs, fmt.Errorf("%s", err))
		return a, nil, errs
	}

	k.RawString = string(s)
	k.SignRequest(true)

	resp, body, errs := k.End()

	if len(errs) > 0 {
		return a, resp, errs
	}

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
	}

	err = json.Unmarshal([]byte(body), &a)

	if err != nil {
		errs = append(errs, err)
		return a, resp, errs
	}

	return a, resp, nil
}

//Delete is a method on an Application which request an Application be deleted in Kumoru.
func (a *Application) Delete() (*Application, *http.Response, []error) {
	k := kumoru.New()

	k.Delete(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, a.UUID))
	k.SignRequest(true)

	resp, _, errs := k.End()

	if len(errs) > 0 {
		return a, resp, errs
	}

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
	}

	return a, resp, nil
}

// Deploy is method on an Application which will cause a deployment in Kumoru.
func (a *Application) Deploy() (*Application, *http.Response, []error) {
	k := kumoru.New()

	k.Post(fmt.Sprintf("%s/v1/applications/%s/deployments/?deployment_token=%s", k.EndPoint.Application, a.UUID, a.DeploymentToken))
	k.SignRequest(true)

	resp, _, errs := k.End()

	if len(errs) > 0 {
		return a, resp, errs
	}

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
	}

	return a, resp, nil
}

// Patch is a method on an application which will modify an existing Application.
func (a *Application) Patch(patchedApplication *Application) (*Application, *http.Response, []error) {
	o, err := json.Marshal(a)
	if err != nil {
		return nil, nil, []error{err}
	}

	p, err := json.Marshal(patchedApplication)
	if err != nil {
		return nil, nil, []error{err}
	}

	patch, err := jsonpatch.CreatePatch([]byte(o), []byte(p))
	if err != nil {
		fmt.Printf("Error creating JSON patch:%v", err)
		return nil, nil, []error{err}
	}

	patchBytes, err := json.Marshal(patch)
	if err != nil {
		return nil, nil, []error{err}
	}
	k := kumoru.New()

	k.Logger.Debugf("Patch string: %s", patchBytes)

	k.Patch(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, a.UUID))
	k.TargetType = "json-patch+json"
	k.RawString = string(string(patchBytes))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if len(errs) > 0 {
		return a, resp, errs
	}

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
	}

	pApp := Application{}
	err = json.Unmarshal([]byte(body), &pApp)

	if err != nil {
		errs = append(errs, err)
		return a, resp, errs
	}

	return &pApp, resp, nil
}

//Show is a method on an Application which retrieves a particular Application from Kumoru.
func (a *Application) Show() (*Application, *http.Response, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, a.UUID))
	k.SignRequest(true)

	resp, body, errs := k.End()

	if len(errs) > 0 {
		return a, resp, errs
	}

	if resp.StatusCode >= 400 {
		errs = append(errs, fmt.Errorf("%s", resp.Status))
	}

	err := json.Unmarshal([]byte(body), &a)

	if err != nil {
		errs = append(errs, fmt.Errorf("%s", err))
		return a, resp, errs
	}

	return a, resp, nil
}

// General functions not explicitly tied to an Application Struct

// List retrieves a list of Applications a role has access to.
func List() (*http.Response, string, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%s/v1/applications/", k.EndPoint.Application))
	k.SignRequest(true)
	return k.End()
}
