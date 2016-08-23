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

package applications

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/client/kumoru/utils"
	"github.com/kumoru/kumoru-sdk-go/pkg/service/application"
	"github.com/ryanuber/columnize"
)

//Archive an Application
func Archive(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		app := application.Application{
			UUID: *uuid,
		}

		_, resp, errs := app.Delete()

		if errs != nil {
			log.Fatalf("Could not archive applications: %s", errs)
		}

		if resp.StatusCode != 202 {
			log.Fatalf("Could not archive applications: %s", resp.Status)
		}

		fmt.Printf("Application %s accepted for archival\n", *uuid)
	}
}

//Create an Application.
func Create(cmd *cli.Cmd) {
	provider := cmd.String(cli.StringArg{
		Name:      "PROVIDER",
		Desc:      "cloud provider to be use",
		HideValue: true,
	})

	region := cmd.String(cli.StringArg{
		Name:      "REGION",
		Desc:      "geographical region to deploy the application",
		HideValue: true,
	})

	image := cmd.String(cli.StringArg{
		Name:      "IMG_URL",
		Desc:      "Image URL",
		HideValue: true,
	})

	name := cmd.String(cli.StringArg{
		Name:      "APP_NAME",
		Desc:      "Application Name",
		HideValue: true,
	})

	certificate := cmd.String(cli.StringOpt{
		Name:      "certificate_file",
		Desc:      "File(PEM encoded) containing the SSL certificate associated with the application",
		HideValue: true,
	})

	certificateChain := cmd.String(cli.StringOpt{
		Name:      "certificate_chain_file",
		Desc:      "File(PEM encoded) contianing the certificate chain associated with the public certificate (optional)",
		HideValue: true,
	})

	envFile := cmd.String(cli.StringOpt{
		Name:      "env_file",
		Desc:      "Environment variables file",
		HideValue: true,
	})

	privateKey := cmd.String(cli.StringOpt{
		Name:      "private_key_file",
		Desc:      "File(PEM encoded) containing the SSL key associated with the public certificate (required if providing a certificate)",
		HideValue: true,
	})

	sslPorts := cmd.Strings(cli.StringsOpt{
		Name:      "ssl_port",
		Desc:      "Port to be associated with the certificate",
		HideValue: true,
	})

	enVars := cmd.Strings(cli.StringsOpt{
		Name:      "e env",
		Desc:      "Environment variable (i.e. MYSQL_PASSWORD=complexpassword",
		HideValue: true,
	})

	rules := cmd.Strings(cli.StringsOpt{
		Name:      "r rule",
		Desc:      "Application Deployment rules",
		HideValue: true,
	})

	ports := cmd.Strings(cli.StringsOpt{
		Name:      "p port",
		Desc:      "Port (non-ssl)",
		HideValue: true,
	})

	labels := cmd.Strings(cli.StringsOpt{
		Name:      "l label",
		Desc:      "Label associated with the application",
		HideValue: true,
	})

	meta := cmd.String(cli.StringOpt{
		Name:      "m metadata",
		Desc:      "Metadata associated with the application being created. Must be JSON formatted.",
		HideValue: true,
	})

	cmd.Action = func() {
		app := application.Application{
			Certificates: readCertificates(certificate, privateKey, certificateChain),
			Environment:  transformEnvironment(envFile, enVars),
			ImageURL:     *image,
			Location: application.Location{
				Provider: *provider,
				Region:   *region,
			},
			Metadata: metaData(*meta, *labels),
			Name:     *name,
			Ports:    *ports,
			Rules:    transformRules(rules),
			SSLPorts: *sslPorts,
		}

		application, resp, errs := app.Create()

		if len(errs) > 0 {
			log.Fatalf("Could not create application: %s", errs[0])
		}

		if resp.StatusCode != 201 {
			log.Fatalf("Could not create application: %s", resp.Status)
		}

		printAppDetail(application)
	}
}

//Deploy an Application
func Deploy(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		app := application.Application{
			UUID: *uuid,
		}

		application, resp, errs := app.Show() // TODO remove this duplication of application.Show() logic

		if errs != nil {
			log.Fatalf("Could not retrieve deployment token: %s", errs)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve deployment token: %s", resp.Status)
		}

		application, resp, errs = application.Deploy()

		if errs != nil {
			log.Fatalf("Could not deploy application: %s", errs)
		}

		if resp.StatusCode != 202 {
			log.Fatalf("Could not deploy application: %s", resp.Status)
		}

		fmt.Printf("Deploying application %s\n", application.UUID)
	}

}

//List all Applications
func List(cmd *cli.Cmd) {
	var a []application.Application

	all := cmd.BoolOpt("a all", false, "List all applications, including archived")

	cmd.Action = func() {
		resp, body, errs := application.List()

		if errs != nil {
			log.Fatalf("Could not retrieve applications: %s", errs)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve applications: %s", resp.Status)
		}

		err := json.Unmarshal([]byte(body), &a)

		if err != nil {
			log.Fatal(err)
		}

		printAppBrief(a, *all)
	}
}

//Show an Application.
func Show(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		app := application.Application{
			UUID: *uuid,
		}

		application, resp, errs := app.Show()

		if errs != nil {
			log.Fatalf("Could not retrieve application: %s", errs)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve application: %s", resp.Status)
		}

		printAppDetail(application)
	}
}

//Patch an Application.
func Patch(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	image := cmd.String(cli.StringOpt{
		Name:      "image_url",
		Desc:      "Image URL",
		HideValue: true,
	})

	name := cmd.String(cli.StringOpt{
		Name:      "name",
		Desc:      "Application Name",
		HideValue: true,
	})

	envFile := cmd.String(cli.StringOpt{
		Name:      "env_file",
		Desc:      "Environment variables file",
		HideValue: true,
	})

	enVars := cmd.Strings(cli.StringsOpt{
		Name:      "e env",
		Desc:      "Environment variable (i.e. MYSQL_PASSWORD=complexpassword",
		HideValue: true,
	})

	rules := cmd.Strings(cli.StringsOpt{
		Name:      "r rule",
		Desc:      "Application Deployment rules",
		HideValue: true,
	})

	labels := cmd.Strings(cli.StringsOpt{
		Name:      "l label",
		Desc:      "Label associated with the aplication",
		HideValue: true,
	})

	meta := cmd.String(cli.StringOpt{
		Name:      "m metadata",
		Desc:      "Metadata associated with the application being created. Must be JSON formatted.",
		HideValue: true,
	})

	cmd.Action = func() {
		originalApp := &application.Application{
			UUID: *uuid,
		}

		originalApp, _, errs := originalApp.Show()

		if len(errs) != 0 {
			log.Fatalf("Unable to retrieve application: %s", errs)
		}

		patchedApp := *originalApp

		if *envFile != "" || len(*enVars) > 0 {
			patchedApp.Environment = transformEnvironment(envFile, enVars)
		} else if *image != "" {
			patchedApp.ImageURL = *image
		} else if *meta != "" || len(*labels) > 0 {
			patchedApp.Metadata = metaData(*meta, *labels)
		} else if *name != "" {
			patchedApp.Name = *name
		} else if len(*rules) > 0 {
			patchedApp.Rules = transformRules(rules)
		}

		pApp, resp, errs := originalApp.Patch(&patchedApp)

		if len(errs) > 0 {
			log.Fatalf("Could not patch application: %s", errs[0])
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not patch application: %s", resp.Status)
		}

		printAppDetail(pApp)
	}
}

func fmtRules(rules map[string]int) string {
	var r string

	for k, v := range rules {
		r += fmt.Sprintf("%s=%v ", k, v)
	}

	return r
}

//metaData combines the provided list of labels with provided arbitary metadata and asserts the result is proper JSON
func metaData(meta string, labels []string) map[string]interface{} {
	js := map[string]interface{}{
		"labels": []string{},
	}

	if len(meta) > 0 {
		err := json.Unmarshal([]byte(meta), &js)
		if err != nil {
			fmt.Println("metadata must be valid JSON:")
			log.Fatal(err)
		}
	}

	if len(labels) > 0 {
		js["labels"] = labels
	}

	return js
}

func transformEnvironment(envFile *string, enVars *[]string) map[string]string {
	var eVars []string
	env := map[string]string{}

	if *envFile != "" {
		eVars = readEnvFile(*envFile)
	} else {
		eVars = *enVars
	}

	for _, v := range eVars {
		e := strings.Split(v, "=")
		env[e[0]] = e[1]
	}

	return env
}

//Transforms a rule string(i.e. "latest=100") to a proper JSON object(i.e. {"latest":100}
func transformRules(rules *[]string) map[string]int {
	rule := map[string]int{}

	for _, v := range *rules {
		r := strings.Split(v, "=")
		weight, err := strconv.Atoi(r[1])
		rule[r[0]] = weight

		if err != nil {
			log.Fatalf("Error converting weight in rule %s to int", v)
		}
	}

	return rule
}

func printAppBrief(a []application.Application, showAll bool) {
	var output []string

	output = append(output, fmt.Sprintf("Name | UUID | Status | Location | Ports | SSL Ports | Rules"))

	for i := 0; i < len(a); i++ {

		if showAll {
			output = append(output, fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s", a[i].Name, a[i].UUID, a[i].Status, a[i].Location, fmtPorts(a[i].Ports), fmtPorts(a[i].SSLPorts), fmtRules(a[i].Rules)))
		} else if strings.ToLower(string(a[i].Status)) != "archived" {
			output = append(output, fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s", a[i].Name, a[i].UUID, a[i].Status, a[i].Location, fmtPorts(a[i].Ports), fmtPorts(a[i].SSLPorts), fmtRules(a[i].Rules)))
		}
	}

	fmt.Println(columnize.SimpleFormat(output))
}

func fmtPorts(ports []string) string {
	if len(ports) > 0 {
		r := ""
		for _, v := range ports {
			r += fmt.Sprintf("%s ", v)
		}
		return r
	}
	return ""
}

func printAppDetail(a *application.Application) {
	var output []string
	var outputEnv []string
	fields := structs.New(a).Fields()

	fmt.Println("\nApplication Details:")

	for _, f := range fields {
		if f.Name() == "Addresses" {
			output = append(output, fmt.Sprintf("%s:\n", f.Name()))
			for _, v := range a.Addresses {
				output = append(output, fmt.Sprintf("……|%s", v))
			}
		} else if f.Name() == "Certificates" {
			output = append(output, fmt.Sprintf("%s:| Use \"--full\" to see certificates", f.Name()))
			output = append(output, fmt.Sprintf("– PrivateKey: |%s\n", a.Certificates.PrivateKey))
		} else if f.Name() == "CreatedAt" {
			output = append(output, fmt.Sprintf("%s: | %s\n", f.Name(), utils.FormatTime(a.CreatedAt+"Z")))
		} else if f.Name() == "CurrentDeployments" {
			output = append(output, fmt.Sprintf("%s:\n", f.Name()))
			for k, v := range a.CurrentDeployments {
				output = append(output, fmt.Sprintf("……|%s: %s", k, v))
			}
		} else if f.Name() == "Environment" {
			outputEnv = append(outputEnv, fmt.Sprintf("%s:\n", f.Name()))
			for k, v := range a.Environment {
				outputEnv = append(outputEnv, fmt.Sprintf("%s=%s", k, v))
			}
		} else if f.Name() == "Location" {
			output = append(output, fmt.Sprintf("%s: |Identifier: %s\t UUID: %s\n", f.Name(), a.Location.Provider, a.Location.Region))
		} else if f.Name() == "Metadata" {
			mdata, _ := json.Marshal(a.Metadata)
			output = append(output, fmt.Sprintf("%s: |%s\n", f.Name(), mdata))
		} else if f.Name() == "Ports" {
			output = append(output, fmt.Sprintf("%s:\n", f.Name()))
			for _, v := range a.Ports {
				output = append(output, fmt.Sprintf("……|%s", v))
			}
		} else if f.Name() == "Rules" {
			output = append(output, fmt.Sprintf("%s:\n", f.Name()))
			for k, v := range a.Rules {
				output = append(output, fmt.Sprintf("……|%s=%v", k, v))
			}
		} else if f.Name() == "SSLPorts" {
			output = append(output, fmt.Sprintf("%s:\n", f.Name()))
			for _, v := range a.SSLPorts {
				output = append(output, fmt.Sprintf("……|%s", v))
			}
		} else if f.Name() == "UpdatedAt" {
			output = append(output, fmt.Sprintf("%s: | %s\n", f.Name(), utils.FormatTime(a.UpdatedAt+"Z")))
		} else {
			output = append(output, fmt.Sprintf("%s: |%v\n", f.Name(), f.Value()))
		}
	}

	fmt.Println(columnize.SimpleFormat(output))
	fmt.Println("\n")
	fmt.Println(columnize.SimpleFormat(outputEnv))
}

func readCertificates(certificate, privateKey, certificateChain *string) application.Certificates {
	var certificates application.Certificates

	if *certificate != "" {
		cert, err := ioutil.ReadFile(*certificate)
		if err != nil {
			log.Fatal(err)
		}
		certificates.Certificate = string(cert)
	}

	if *privateKey != "" {
		key, err := ioutil.ReadFile(*privateKey)
		if err != nil {
			log.Fatal(err)
		}
		certificates.PrivateKey = string(key)
	}

	if *certificateChain != "" {
		chain, err := ioutil.ReadFile(*certificateChain)
		if err != nil {
			log.Fatal(err)
		}
		certificates.CertificateChain = string(chain)
	}

	return certificates
}

func readEnvFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	x := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x = append(x, scanner.Text())
	}

	return x
}
