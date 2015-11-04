package application

import (
	"fmt"

	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

func Create(name, image string, envVars, rules, ports []string) {

	k := kumoru.New()

	k = k.Post(fmt.Sprintf("%s/v1/applications/", k.EndPoint.Application))

	k = k.Send(fmt.Sprintf("name=%s&", name))
	k = k.Send(fmt.Sprintf("image_url=%s&", image))

	for _, envVar := range envVars {
		k = k.Send(fmt.Sprintf("environment=%s&", envVar))
	}

	for _, port := range ports {
		k = k.Send(fmt.Sprintf("port=%s&", port))
	}

	for _, rule := range rules {
		k = k.Send(fmt.Sprintf("rule=%s&", rule))
	}

	resp, body, errs := k.SignRequest(true).
		End()

	k.Logger.Info("resp: ", resp)
	k.Logger.Info("body: ", body)
	k.Logger.Info("errs: ", errs)
}

func List() {
	k := kumoru.New()

	resp, body, errs := k.Get(fmt.Sprintf("%s/v1/applications/", k.EndPoint.Application)).
		SignRequest(true).
		End()

	k.Logger.Info("resp: ", resp)
	k.Logger.Info("body: ", body)
	k.Logger.Info("errs: ", errs)
}

func Show(uuid string) {
	k := kumoru.New()

	resp, body, errs := k.Get(fmt.Sprintf("%s/v1/applications/%s", k.EndPoint.Application, uuid)).
		SignRequest(true).
		End()

	k.Logger.Info("resp: ", resp)
	k.Logger.Info("body: ", body)
	k.Logger.Info("errs: ", errs)
}

func ApplicationDeploy(uuid string) {
	k := kumoru.New()

	resp, body, errs := k.Post(fmt.Sprintf("%s/v1/applications/%s/deployments/", k.EndPoint.Application, uuid)).
		SignRequest(true).
		End()

	k.Logger.Info("resp: ", resp)
	k.Logger.Info("body: ", body)
	k.Logger.Info("errs: ", errs)
}

func ApplicationDelete(uuid string) {
	k := kumoru.New()

	resp, body, errs := k.Post(fmt.Sprintf("%s/v1/applications/%s/deployments/", k.EndPoint.Application, uuid)).
		SignRequest(true).
		End()

	k.Logger.Info("resp: ", resp)
	k.Logger.Info("body: ", body)
	k.Logger.Info("errs: ", errs)
}
