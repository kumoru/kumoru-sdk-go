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

package resources

import (
	"fmt"
	"net/http"

	"github.com/kumoru/kumoru-sdk-go/pkg/kumoru"
)

//Resource respresents a resource model in authorization-manager API.
type Resource struct {
	CreatedAt  string `json:"created_at"`
	Type       string `json:"type"`
	UpdatedAt  string `json:"updated_at"`
	Identifier string `json:"identifier"`
	Context    string `json:"context"`
}

// Find resources that are accesible to the requester
func Find(rType, action, identifier string, wrappedRequest *http.Request) (*http.Response, string, []error) {
	params := "select_by="
	params += fmt.Sprintf("type=%s,", rType)
	params += fmt.Sprintf("action=%s", action)
	if identifier != "" {
		params += fmt.Sprintf(",uuid=%s,", identifier)
	}

	fmt.Println("Params: ", params)

	k := kumoru.New()
	k.Get(fmt.Sprintf("%s/v1/resources/", k.EndPoint.Authorization))
	k.Query(params)
	if wrappedRequest != nil {
		k.ProxyRequest(wrappedRequest)
		if wrappedRequest.Header.Get("X-Kumoru-Context") != "" {
			k.SetHeader("X-Kumoru-Context", wrappedRequest.Header.Get("X-Kumoru-Context"))
		}
	}
	k.SignRequest(true)

	return k.End()
}
