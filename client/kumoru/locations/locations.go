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

package locations

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/pkg/service/location"
	"github.com/ryanuber/columnize"
)

//Add a location in a given region for the particular provider
func Add(cmd *cli.Cmd) {
	provider := cmd.String(cli.StringArg{
		Name:      "PROVIDER",
		Desc:      "Cloud provider(i.e. amazon)",
		HideValue: true,
	})

	identifier := cmd.String(cli.StringArg{
		Name:      "IDENTIFIER",
		Desc:      "Cloud provider specific region/zone/etc identifier (i.e. us-east-1)",
		HideValue: true,
	})

	cmd.Action = func() {
		l := location.Location{
			Provider: *provider,
			Region:   *identifier,
		}

		body, errs := l.Create()

		if len(errs) > 0 {
			log.Fatalf("Could not add new location: %s", errs)
		}

		err := json.Unmarshal([]byte(body), &l)

		if err != nil {
			log.Fatal(err)
		}

		PrintLocationBrief([]location.Location{l})
	}
}

//Delete a location in a given region for the particular provider
func Delete(cmd *cli.Cmd) {
	provider := cmd.String(cli.StringArg{
		Name:      "PROVIDER",
		Desc:      "Cloud provider(i.e. amazon)",
		HideValue: true,
	})

	identifier := cmd.String(cli.StringArg{
		Name:      "IDENTIFIER",
		Desc:      "Cloud provider specific region/zone/etc identifier (i.e. us-east-1)",
		HideValue: true,
	})

	cmd.Action = func() {
		l := location.Location{
			Provider: *provider,
			Region:   *identifier,
		}

		//TODO determine if there are any applications in the location and prompt user to remove them
		errs := l.Delete()

		if len(errs) > 0 {
			log.Fatalf("Could not delete location: %s", errs)
		}

		fmt.Printf("Deleting location %s-%s", *provider, *identifier)
	}
}

//List all available Locations and optionally apply a filter
func List(cmd *cli.Cmd) {
	provider := cmd.String(cli.StringOpt{
		Name:      "provider",
		Desc:      "Cloud provider(i.e. amazon)",
		HideValue: true,
	})

	identifier := cmd.String(cli.StringOpt{
		Name:      "identifier",
		Desc:      "Cloud provider specific region/zone/etc identifier (i.e. us-east-1)",
		HideValue: true,
	})

	cmd.Action = func() {
		l := location.Location{
			Provider: *provider,
			Region:   *identifier,
		}

		body, errs := l.Find()

		if len(errs) > 0 {
			log.Fatalf("Could not retrieve locations: %s", errs[0])
		}

		locations := &[]location.Location{}
		err := json.Unmarshal([]byte(body), locations)

		if err != nil {
			log.Fatal(err)
		}

		PrintLocationBrief(*locations)
	}
}

//PrintLocationBrief outputs a listing of locations with minimal details
func PrintLocationBrief(l []location.Location) {
	var output []string

	output = append(output, fmt.Sprintf("Provider | Region"))

	for i := 0; i < len(l); i++ {
		output = append(output, fmt.Sprintf("%s | %s", l[i].Provider, l[i].Region))
	}

	fmt.Println(columnize.SimpleFormat(output))
}
