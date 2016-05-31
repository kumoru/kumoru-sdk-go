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

//Package appllication provides an Application type and pertinent methods for it.
package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kumoru/kumoru-sdk-go/pkg/kumoru"
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
	Location           map[string]string      `json:"location"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
	Name               string                 `json:"name"`
	Ports              []string               `json:"ports,omitempty"`
	Rules              map[string]int         `json:"rules,omitempty"`
	SSLPorts           []string               `json:"ssl_ports,omitempty"`
	Status             string                 `json:"status,omitempty"`
	UpdatedAt          string                 `json:"updated_at,omitempty"`
	URL                string                 `json:"url,omitempty"`
	UUID               string                 `json:"uuid,omitempty"`
	ApiVersion         string                 `json:"api_version,omitempty"`
	Certificates       Certificates           `json:"certificates,omitempty"`
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

// Patch is a method on an application which will modify an Application.
func (a *Application) Patch(certificates, name, image, metaData string, envVars, rules, ports, sslPorts []string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Patch(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, a.UUID))
	k.Send(genParameters(certificates, name, image, metaData, envVars, rules, ports, sslPorts))
	k.SignRequest(true)
	return k.End()
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

//genParameters decides which query strings to append to the URI.
//TODO remove once PATCH has a json interface
func genParameters(certificates, name, image, metaData string, envVars, rules, ports, sslPorts []string) string {
	var params string

	if certificates != "" {
		params += fmt.Sprintf("certificates=%s&", url.QueryEscape(certificates))
	}

	if name != "" {
		params += fmt.Sprintf("name=%s&", url.QueryEscape(name))
	}

	if image != "" {
		params += fmt.Sprintf("image_url=%s&", url.QueryEscape(image))
	}

	if metaData != "" {
		params += fmt.Sprintf("metadata=%s&", url.QueryEscape(metaData))
	}

	for _, envVar := range envVars {
		params += fmt.Sprintf("environment=%s&", url.QueryEscape(envVar))
	}

	for _, port := range ports {
		params += fmt.Sprintf("ports=%s&", url.QueryEscape(port))
	}

	for _, sslport := range sslPorts {
		params += fmt.Sprintf("ssl_ports=%s&", url.QueryEscape(sslport))
	}

	for _, rule := range rules {
		params += fmt.Sprintf("rules=%s&", url.QueryEscape(rule))
	}

	return params
}
