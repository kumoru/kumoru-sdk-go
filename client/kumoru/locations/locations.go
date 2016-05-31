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
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/pkg/service/pools"
	"github.com/ryanuber/columnize"
)

func Add(cmd *cli.Cmd) {

	identifier := cmd.String(cli.StringArg{
		Name:      "LOCATION",
		Desc:      "location to be added to your current role(i.e. us-east-1)",
		HideValue: true,
	})

	cmd.Action = func() {
		l := pools.Location{}
		l.Identifier = *identifier

		location, resp, errs := l.Create()

		if len(errs) > 0 {
			log.Fatalf("Could not add new location: %s", errs)
		}

		if resp.StatusCode != 201 {
			log.Fatalf("Cloud not add new location: %s", resp.Status)
		}

		PrintLocationBrief([]pools.Location{*location}, false)
	}
}

func Archive(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Region UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		var l *pools.Location
		l, resp, errs := l.Delete(*uuid)

		if len(errs) > 0 {
			log.Fatalf("Could not archive location: %s", errs)
		}

		if resp.StatusCode != 202 {
			log.Fatalf("Could not archive location: %s", resp.Status)
		}

		fmt.Sprintf("Location %s accepted for archival\n", *uuid)
	}
}

func List(cmd *cli.Cmd) {
	all := cmd.BoolOpt("a all", false, "List all locations, including archived")

	cmd.Action = func() {
		l := pools.Location{}
		locations, resp, errs := l.List()

		if len(errs) > 0 {
			log.Fatalf("Could not retrieve locations: %s", errs[0])
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Cloud not retrieve locations: %s", resp.Status)
		}

		PrintLocationBrief(*locations, *all)
	}
}

func PrintLocationBrief(l []pools.Location, showAll bool) {
	var output []string

	output = append(output, fmt.Sprintf("Location | Provider | UUID | Status | Aggregrate Resources"))

	for i := 0; i < len(l); i++ {
		if showAll {
			output = append(output, fmt.Sprintf("%s | %s | %s| %s| %v vCPU, %vGB RAM", l[i].Identifier, l[i].Provider, l[i].Uuid, l[i].Status, l[i].AggregateResources["cpu"], l[i].AggregateResources["ram"]))
		} else if strings.ToLower(string(l[i].Status)) != "archived" {
			output = append(output, fmt.Sprintf("%s | %s | %s| %s| %v vCPU, %vGB RAM", l[i].Identifier, l[i].Provider, l[i].Uuid, l[i].Status, l[i].AggregateResources["cpu"], l[i].AggregateResources["ram"]))
		}
	}

	fmt.Println(columnize.SimpleFormat(output))
}
