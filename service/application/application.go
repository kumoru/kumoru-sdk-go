package application

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

func Create(name, image string, envVars, rules, ports []string) (*http.Response, string, []error) {

	k := kumoru.New()

	k = k.Post(fmt.Sprintf("%s/v1/applications/", k.EndPoint.Application))

	k = k.Send(genParameters(name, image, envVars, rules, ports))

	return k.SignRequest(true).
		End()
}

func List() (*http.Response, string, []error) {
	k := kumoru.New()

	return k.Get(fmt.Sprintf("%s/v1/applications/", k.EndPoint.Application)).
		SignRequest(true).
		End()
}

func Show(uuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	return k.Get(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, uuid)).
		SignRequest(true).
		End()
}

func Deploy(uuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	return k.Post(fmt.Sprintf("%s/v1/applications/%s/deployments/", k.EndPoint.Application, uuid)).
		SignRequest(true).
		End()
}

func Delete(uuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	resp, body, errs := k.Post(fmt.Sprintf("%s/v1/applications/%s/deployments/", k.EndPoint.Application, uuid)).
		SignRequest(true).
		End()

	return resp, body, errs
}

func Patch(uuid, name, image string, envVars, rules, ports []string) (*http.Response, string, []error) {

	k := kumoru.New()

	k = k.Patch(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, uuid))

	k = k.Send(genParameters(name, image, envVars, rules, ports))

	return k.SignRequest(true).
		End()
}

func genParameters(name, image string, envVars, rules, ports []string) string {
	var params string

	if name != "" {
		params += fmt.Sprintf("name=%s&", name)
	}

	if image != "" {
		params += fmt.Sprintf("image_url=%s&", image)
	}

	for _, envVar := range envVars {
		params += fmt.Sprintf("environment=%s&", envVar)
	}

	for _, port := range ports {
		params += fmt.Sprintf("ports=%s&", port)
	}

	for _, rule := range rules {
		params += fmt.Sprintf("rule=%s&", rule)
	}

	return url.QueryEscape(params)
}
