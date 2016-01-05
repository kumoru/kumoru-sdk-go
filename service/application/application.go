package application

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

func Create(poolUuid, certificates, name, image, providerCredentials, metaData string, envVars, rules, ports, sslPorts []string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Post(fmt.Sprintf("%s/v1/applications/", k.EndPoint.Application))

	k.Send(genParameters(certificates, name, image, providerCredentials, metaData, envVars, rules, ports, sslPorts))

	k.Send(fmt.Sprintf("pool_uuid=%s&", url.QueryEscape(poolUuid)))

	k.SignRequest(true)

	return k.End()
}

func List() (*http.Response, string, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%s/v1/applications/", k.EndPoint.Application))
	k.SignRequest(true)
	return k.End()
}

func Show(applicationUuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Get(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, applicationUuid))
	k.SignRequest(true)
	return k.End()
}

func Deploy(applicationUuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Post(fmt.Sprintf("%s/v1/applications/%s/deployments/", k.EndPoint.Application, applicationUuid))
	k.SignRequest(true)
	return k.End()
}

func Delete(applicationUuid string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Delete(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, applicationUuid))
	k.SignRequest(true)
	return k.End()

}

func Patch(applicationUuid, certificates, name, image, providerCredentials, metaData string, envVars, rules, ports, sslPorts []string) (*http.Response, string, []error) {
	k := kumoru.New()

	k.Patch(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, applicationUuid))

	k.Send(genParameters(certificates, name, image, providerCredentials, metaData, envVars, rules, ports, sslPorts))

	k.SignRequest(true)
	return k.End()
}

func genParameters(certificates, name, image, providerCredentials, metaData string, envVars, rules, ports, sslPorts []string) string {
	var params string

	if certificates != "" {
		params += fmt.Sprintf("certificates=%s&", url.QueryEscape(certificates))
	}

	if name != "" {
		params += fmt.Sprintf("name=%s&", url.QueryEscape(name))
	}

	if image != "" {
		params += fmt.Sprintf("image_url=%s&", url.QueryEscape(image))
	}

	if providerCredentials != "" {
		params += fmt.Sprintf("provider_credentials=%s&", url.QueryEscape(providerCredentials))
	}

	if metaData != "" {
		params += fmt.Sprintf("metadata=%s&", url.QueryEscape(metaData))
	}

	for _, envVar := range envVars {
		params += fmt.Sprintf("environment=%s&", url.QueryEscape(envVar))
	}

	for _, port := range ports {
		params += fmt.Sprintf("ports=%s&", url.QueryEscape(port))
	}

	for _, sslport := range sslPorts {
		params += fmt.Sprintf("ssl_ports=%s&", url.QueryEscape(sslport))
	}

	for _, rule := range rules {
		params += fmt.Sprintf("rules=%s&", url.QueryEscape(rule))
	}

	return params
}
