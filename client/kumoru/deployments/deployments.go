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

package deployments

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/client/kumoru/utils"
	"github.com/kumoru/kumoru-sdk-go/pkg/service/application/deployments"
	"github.com/ryanuber/columnize"
)

func List(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		d := deployments.Deployment{}
		deployments, resp, errs := d.List(*uuid)

		if len(errs) > 0 {
			log.Fatalf("Could not retrieve deployments: %s", errs[0])
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve deployments: %s", resp.Status)
		}

		printDeploymentsBrief(*deployments)
	}
}

func Show(cmd *cli.Cmd) {
	applicationUuid := cmd.String(cli.StringArg{
		Name:      "APPLICATION_UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	uuid := cmd.String(cli.StringArg{
		Name:      "DEPLOYMENT_UUID",
		Desc:      "Deployment UUID",
		HideValue: true,
	})
	cmd.Action = func() {
		d := deployments.Deployment{}
		deployment, resp, errs := d.Show(*applicationUuid, *uuid)

		if len(errs) > 0 {
			log.Fatalf("Could not retrieve deployment: %s", errs[0])
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve deployment: %s", resp.Status)
		}

		printDeploymentDetail(*deployment)
	}
}

func printDeploymentsBrief(d []deployments.Deployment) {
	var output []string

	output = append(output, fmt.Sprintf("UUID | Created At | Image Tag | Image Id"))

	for i := 0; i < len(d); i++ {
		output = append(output, fmt.Sprintf("%s | %s | %s | %s", d[i].Uuid, d[i].CreatedAt, d[i].ImageTag, d[i].ImageId))
	}

	fmt.Println(columnize.SimpleFormat(output))
}

func printDeploymentDetail(d deployments.Deployment) {
	var output []string
	fields := structs.New(d).Fields()

	fmt.Println("\nDeployment Details:\n")

	for _, f := range fields {
		if f.Name() == "Metadata" {
			mdata, _ := json.Marshal(d.Metadata)
			output = append(output, fmt.Sprintf("%s: |%s\n", f.Name(), mdata))
		} else if f.Name() == "CreatedAt" {
			output = append(output, fmt.Sprintf("%s: | %s\n", f.Name(), utils.FormatTime(d.CreatedAt+"Z")))
		} else {
			output = append(output, fmt.Sprintf("%s: |%v\n", f.Name(), f.Value()))
		}
	}

	fmt.Println(columnize.SimpleFormat(output))
}
