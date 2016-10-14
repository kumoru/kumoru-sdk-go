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

package kumoru

import (
	"os"

	"github.com/go-ini/ini"
)

// Endpoints struct for all api services
type Endpoints struct {
	Application   string
	Authorization string
	Location      string
}

// LoadEndpoints returns and Endpoints struct by reading them from a file or from defaults
func LoadEndpoints(filename string, section string) Endpoints {
	config, err := ini.Load(filename)

	appManagerURL := "https://application.api.kumoru.io"
	if os.Getenv("APPLICATION_MANAGER_URL") != "" {
		appManagerURL = os.Getenv("APPLICATION_MANAGER_URL")
	}

	authManagerURL := "https://authorization.api.kumoru.io"
	if os.Getenv("AUTHORIZATION_MANAGER_URL") != "" {
		authManagerURL = os.Getenv("AUTHORIZATION_MANAGER_URL")
	}

	locationManagerURL := "https://location.api.kumoru.io"
	if os.Getenv("LOCATION_MANAGER_URL") != "" {
		locationManagerURL = os.Getenv("LOCATION_MANAGER_URL")
	}

	if err != nil {
		return Endpoints{
			Application:   appManagerURL,
			Authorization: authManagerURL,
			Location:      locationManagerURL,
		}
	}

	iniEndpoints, err := config.GetSection(section)
	if err != nil {
		return Endpoints{
			Application:   appManagerURL,
			Authorization: authManagerURL,
			Location:      locationManagerURL,
		}
	}

	return Endpoints{
		Application:   iniEndpoints.Key("kumoru_application_api").String(),
		Authorization: iniEndpoints.Key("kumoru_authorization_api").String(),
		Location:      iniEndpoints.Key("kumoru_location_api").String(),
	}
}
