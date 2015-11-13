package application

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

func Create(poolUuid, name, image string, envVars, providerCredentials, rules, ports []string) (*http.Response, string, []error) {

	k := kumoru.New()

	k = k.Post(fmt.Sprintf("%s/v1/applications/", k.EndPoint.Application))

	k = k.Send(genParameters(name, image, envVars, providerCredentials, rules, ports))

	k = k.Send(fmt.Sprintf("pool_uuid=%s&", url.QueryEscape(poolUuid)))

	return k.SignRequest(true).
		End()
}

func List() (*http.Response, string, []error) {
	k := kumoru.New()

	return k.Get(fmt.Sprintf("%s/v1/applications/", k.EndPoint.Application)).
		SignRequest(true).
		End()
}

func Show(applicationUuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	return k.Get(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, applicationUuid)).
		SignRequest(true).
		End()
}

func Deploy(applicationUuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	return k.Post(fmt.Sprintf("%s/v1/applications/%s/deployments/", k.EndPoint.Application, applicationUuid)).
		SignRequest(true).
		End()
}

func Delete(applicationUuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	resp, body, errs := k.Delete(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, applicationUuid)).
		SignRequest(true).
		End()

	return resp, body, errs
}

func Patch(applicationUuid, name, image string, envVars, providerCredentials, rules, ports []string) (*http.Response, string, []error) {

	k := kumoru.New()

	k = k.Patch(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, applicationUuid))

	k = k.Send(genParameters(name, image, envVars, providerCredentials, rules, ports))

	return k.SignRequest(true).
		End()
}

func genParameters(name, image string, envVars, providerCredentials, rules, ports []string) string {
	var params string

	if name != "" {
		params += fmt.Sprintf("name=%s&", url.QueryEscape(name))
	}

	if image != "" {
		params += fmt.Sprintf("image_url=%s&", url.QueryEscape(image))
	}

	for _, providerCredentials := range providerCredentials {
		params += fmt.Sprintf("provider_credentials=%s&", url.QueryEscape(providerCredentials))
	}

	for _, envVar := range envVars {
		params += fmt.Sprintf("environment=%s&", url.QueryEscape(envVar))
	}

	for _, port := range ports {
		params += fmt.Sprintf("ports=%s&", url.QueryEscape(port))
	}

	for _, rule := range rules {
		params += fmt.Sprintf("rule=%s&", url.QueryEscape(rule))
	}

	return params
}
