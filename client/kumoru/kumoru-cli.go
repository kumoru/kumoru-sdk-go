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

package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/jawher/mow.cli"

	"github.com/kumoru/kumoru-sdk-go/client/kumoru/accounts"
	"github.com/kumoru/kumoru-sdk-go/client/kumoru/applications"
	"github.com/kumoru/kumoru-sdk-go/client/kumoru/deployments"
	"github.com/kumoru/kumoru-sdk-go/client/kumoru/locations"
	"github.com/kumoru/kumoru-sdk-go/client/kumoru/secrets"
	"github.com/kumoru/kumoru-sdk-go/client/kumoru/tokens"
)

func init() {
	// Initialize Logging level to WARN
	// Need to change this to be configurable
	log.SetLevel(log.DebugLevel)
	///log.SetLevel(log.WarnLevel)
	log.SetOutput(os.Stderr)
}

var Version = "0.1.7"
var GitVersion = "No Version Provided"
var BuildStamp = "No Build Stamp Provided"

func main() {

	BuildVersion := fmt.Sprintf("Version: %s \nGit Commit Hash: %s \nUTC Build Time: %s", Version, GitVersion, BuildStamp)

	app := cli.App("kumoru", "Utility to interact with Kumoru services.")

	app.Version("v version", BuildVersion)

	app.Command("login", "Login action", tokens.Create)

	app.Command("accounts", "Account actions", func(act *cli.Cmd) {
		act.Command("create", "Create an account ", accounts.Create)
		act.Command("reset", "Reset password for account", accounts.ResetPassword)
		act.Command("show", "Show account information", accounts.Show)
	})

	app.Command("applications", "Application actions", func(apps *cli.Cmd) {
		apps.Command("archive", "Archive an application", applications.Archive)
		apps.Command("create", "Create an application", applications.Create)
		apps.Command("deploy", "Deploy an application", applications.Deploy)
		apps.Command("list", "List all applications", applications.List)
		apps.Command("patch", "Update an application", applications.Patch)
		apps.Command("show", "Show application information", applications.Show)
	})

	app.Command("deployments", "Deployment actions", func(apps *cli.Cmd) {
		apps.Command("list", "List all deployments", deployments.List)
		apps.Command("show", "Show deployment information", deployments.Show)
	})

	app.Command("locations", "Location actions", func(location *cli.Cmd) {
		location.Command("add", "Add location to current role", locations.Add)
		location.Command("archive", "Archive specific location (WARNING: destructive)", locations.Archive)
		location.Command("list", "List all locations", locations.List)
	})

	app.Command("secrets", "Secrets actions", func(sec *cli.Cmd) {
		sec.Command("create", "Create secret", secrets.Create)
		sec.Command("list", "List secrets", secrets.List)
		sec.Command("show", "Show secret", secrets.Show)
	})

	app.Command("tokens", "Token actions", func(tkn *cli.Cmd) {
		tkn.Command("create", "Create a pair of tokens", tokens.Create)
	})

	app.Run(os.Args)
}
