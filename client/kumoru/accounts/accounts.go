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

package accounts

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-cli/utils"
	"github.com/kumoru/kumoru-sdk-go/pkg/service/authorization"
	"github.com/ryanuber/columnize"
	"golang.org/x/crypto/ssh/terminal"
)

//Create an Account.
func Create(cmd *cli.Cmd) {
	email := cmd.String(cli.StringArg{
		Name:      "EMAIL",
		Desc:      "email address",
		HideValue: true,
	})

	fName := cmd.String(cli.StringOpt{
		Name:      "f first-name",
		Desc:      "Given Name",
		HideValue: true,
	})

	lName := cmd.String(cli.StringOpt{
		Name:      "l last-name",
		Desc:      "Last Name",
		HideValue: true,
	})

	password := cmd.String(cli.StringOpt{
		Name:      "p password",
		Desc:      "Password",
		HideValue: true,
	})

	cmd.Action = func() {

		if *password == "" {
			*password = passwordPrompt()
			fmt.Println("\n")
		}

		l := authorization.Account{
			Email:     *email,
			GivenName: *fName,
			Surname:   *lName,
		}

		account, resp, errs := l.CreateAcct(*password)

		if len(errs) > 0 {
			log.Fatalf("Could not create account: %s", errs[0])
		}

		if resp.StatusCode != 201 {
			log.Fatalf("Could not create account: %s", resp.Status)
		}

		printAccountDetail(account)
	}
}

//Show displays the details of an Account.
func Show(cmd *cli.Cmd) {
	email := cmd.String(cli.StringArg{
		Name:      "EMAIL",
		Desc:      "email address",
		HideValue: true,
	})

	cmd.Action = func() {
		a := authorization.Account{
			Email: *email,
		}

		account, resp, errs := a.ShowAcct()

		if len(errs) > 0 {
			log.Fatalf("Could not retrieve account: %s", errs[0])
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve account: %s", resp.Status)
		}

		printAccountDetail(account)
	}
}

//ResetPassword requests the password reset operation to start.
func ResetPassword(cmd *cli.Cmd) {
	email := cmd.String(cli.StringArg{
		Name:      "EMAIL",
		Desc:      "email address",
		HideValue: true,
	})

	cmd.Action = func() {
		a := authorization.Account{
			Email: *email,
		}

		account, resp, errs := a.ResetPassword()

		if len(errs) > 0 {
			log.Fatalf("Could not retrieve account: %s", errs[0])
		}

		if resp.StatusCode != 204 {
			log.Fatalf("Could not reset account password: %s", resp.Status)
		}

		fmt.Println("Password reset instructions sent to:", account.Email)
	}
}

//Private functions

//Prints details of an Account. Fields are formatted to be more human readable.
func printAccountDetail(a *authorization.Account) {
	var output []string
	fields := structs.New(a).Fields()

	fmt.Println("\nAccount Details:\n")

	for _, f := range fields {
		if f.Name() == "CreatedAt" {
			output = append(output, fmt.Sprintf("%s: | %s\n", f.Name(), utils.FormatTime(a.CreatedAt+"Z")))
		} else if f.Name() == "UpdatedAt" {
			output = append(output, fmt.Sprintf("%s: | %s\n", f.Name(), utils.FormatTime(a.UpdatedAt+"Z")))
		} else {
			output = append(output, fmt.Sprintf("%s: |%s\n", f.Name(), f.Value()))
		}
	}

	fmt.Println(columnize.SimpleFormat(output))
}

//passwordPrompt prompts user to enter password twice and that the two entries are equivalent.
func passwordPrompt() string {
	fmt.Print("Enter password: ")
	password, errs := terminal.ReadPassword(0)

	if errs != nil {
		fmt.Println("\nCould not read password:")
		log.Fatal(errs)
		os.Exit(1)
	}

	fmt.Print("\nConfirm password: ")
	passwordConfirm, errs := terminal.ReadPassword(0)

	if errs != nil {
		fmt.Println("\nCould Not read password.")
		log.Fatal(errs)
		os.Exit(1)
	}

	if reflect.DeepEqual(password, passwordConfirm) == false {
		fmt.Println("\n")
		log.Fatal("Passwords do not match")
	}

	return strings.TrimSpace(string(password))
}
