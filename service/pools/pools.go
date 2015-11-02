package pools

import (
	"fmt"

	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

func Create(location string, credentials string) {
	k := kumoru.New()

	if location == "" || credentials == "" {
		k.Logger.Fatal("Please provide a location or credentials")
	}

	if k.Tokens.Public == "" || k.Tokens.Private == "" {
		k.Logger.Fatal("Update your config with a set of tokens")
	}

	resp, _, errs := k.Post(fmt.Sprintf("%v/v1/pools/", k.EndPoint.Pool)).
		Send(fmt.Sprintf("location=%s&credentials=%s", location, credentials)).
		SignRequest(true).
		SetDebug(true).
		End()

	k.Logger.Warn("k.Data = ", k.Data)
	k.Logger.Warn("k.RawString = ", k.RawString)

	if errs != nil {
		k.Logger.Fatal("Fatal error Error: ", errs)
	} else {
		if resp.StatusCode == 201 {
			k.Logger.Warn("Pool Created: ", resp.Header["Location"])
			fmt.Println(resp.Status)
		} else {
			k.Logger.Debug("resp: ", resp)
			k.Logger.Warn("Unable to create pool: ", resp.Status)
		}
	}
}

func List(bootstrap bool) {

	k := kumoru.New()

	resp, body, errs := k.Get(fmt.Sprintf("%v/v1/pools/", k.EndPoint.Pool)).
		SignRequest(true).
		SetDebug(true).
		End()

	k.Logger.Info("body: ", body)
	k.Logger.Info("resp: ", resp)
	if errs != nil {
		k.Logger.Fatal("Fatal error Error: ", errs)
	} else {
		if resp.StatusCode == 200 {
			k.Logger.Info("Pool list: ", body)
			k.Logger.Warn("Pool list: ", body)
			fmt.Println(body)
		}
	}

}
