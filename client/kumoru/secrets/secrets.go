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

package secrets

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/client/kumoru/utils"
	"github.com/kumoru/kumoru-sdk-go/pkg/service/authorization/secrets"
	"github.com/ryanuber/columnize"
)

func Create(cmd *cli.Cmd) {
	value := cmd.String(cli.StringArg{
		Name:      "VALUE",
		Desc:      "Value to be stored as secret",
		HideValue: true,
	})

	labels := cmd.Strings(cli.StringsOpt{
		Name:      "l label",
		Desc:      "Label to attach to Secret. This option may be included more than once.",
		HideValue: true,
	})

	cmd.Action = func() {
		s := secrets.Secret{
			Value:  *value,
			Labels: *labels,
		}

		secret, resp, errs := s.Create()

		if len(errs) > 0 {
			log.Fatalf("Could not create secret: %s", errs[0])
		}

		if resp.StatusCode != 201 {
			log.Fatalf("Could not create secret: %s", resp.Status)
		}

		printSecretDetail(secret)
	}
}

func List(cmd *cli.Cmd) {
	cmd.Action = func() {
		secrets, resp, errs := secrets.List()

		if len(errs) > 0 {
			log.Fatalf("Could not retrieve secret: %s", errs[0])
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve secret: %s", resp.Status)
		}

		printSecretBrief(secrets)
	}
}

func Show(cmd *cli.Cmd) {
	secretUuid := cmd.String(cli.StringArg{
		Name:      "SECRET_UUID",
		Desc:      "UUID of secret to retrieve",
		HideValue: true,
	})

	cmd.Action = func() {
		s := secrets.Secret{}
		secret, resp, errs := s.Show(secretUuid)

		if len(errs) > 0 {
			log.Fatalf("Could not retrieve secret: %s", errs[0])
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve secret: %s", resp.Status)
		}

		printSecretDetail(secret)
	}
}

func printSecretBrief(apps []*secrets.Secret) {
	var output []string

	output = append(output, fmt.Sprintf("UUID | Created At | Labels"))

	for i := 0; i < len(apps); i++ {
		output = append(output, fmt.Sprintf("%s | %s | %s", apps[i].Uuid, utils.FormatTime(apps[i].CreatedAt+"Z"), apps[i].Labels))
	}

	fmt.Println(columnize.SimpleFormat(output))
}

func printSecretDetail(s *secrets.Secret) {
	var output []string
	fields := structs.New(s).Fields()

	fmt.Println("\nSecret Details:\n")

	for _, f := range fields {
		if f.Name() == "CreatedAt" {
			output = append(output, fmt.Sprintf("%s: | %s\n", f.Name(), utils.FormatTime(s.CreatedAt+"Z")))
		} else if f.Name() == "UpdatedAt" {
			output = append(output, fmt.Sprintf("%s: | %s\n", f.Name(), utils.FormatTime(s.UpdatedAt+"Z")))
		} else {
			output = append(output, fmt.Sprintf("%s: |%v\n", f.Name(), f.Value()))
		}
	}

	fmt.Println(columnize.SimpleFormat(output))
}
