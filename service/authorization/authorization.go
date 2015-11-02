package authorization

import (
	"fmt"

	"code.google.com/p/go-uuid/uuid"
	"github.com/kumoru/kumoru-sdk-go/kumoru"
)

func GetTokens() (token, secret string) {

	k := kumoru.New()

	token = uuid.New()

	resp, body, errs := k.Put(fmt.Sprintf("%v/v1/tokens/%v", k.EndPoint.Authorization, token)).
		End()

	if errs != nil {
		k.Logger.Fatal("Could not retrieve new tokens")
	}

	if resp.StatusCode != 201 {
		k.Logger.Fatal(fmt.Sprintf("Could not retrieve token, server responded: %v", resp.Status))
	}

	return token, body
}

func CreateAcct(username, given_name, surname, password string) {

	k := kumoru.New()

	resp, body, errs := k.Put(fmt.Sprintf("%s/v1/accounts/%s", k.EndPoint.Authorization, username)).
		Send(fmt.Sprintf("given_name=%s&surname=%s&password=%s", given_name, surname, password)).
		End()

	if errs != nil {
		k.Logger.Fatal("Could not retrieve new tokens")
	}

	switch resp.StatusCode {
	case 200:
		fmt.Println("Account created successfully")
		k.Logger.Info(fmt.Sprintf("Response Status: %v", resp.Status))
		k.Logger.Info(fmt.Sprintf("Response: %v", resp))
		k.Logger.Info(fmt.Sprintf("body: %v", body))
	case 409:
		fmt.Println("Account already exists.")
		k.Logger.Info(fmt.Sprintf("Response Status: %v", resp.Status))
		k.Logger.Info(fmt.Sprintf("Response: %v", resp))
		k.Logger.Info(fmt.Sprintf("body: %v", body))
	default:
		k.Logger.Info(fmt.Sprintf("Response Status: %v", resp.Status))
		k.Logger.Info(fmt.Sprintf("Response: %v", resp))
		k.Logger.Info(fmt.Sprintf("body: %v", body))
	}
}
func ShowAcct(username string) {

	k := kumoru.New()

	resp, body, errs := k.Get(fmt.Sprintf("%v/v1/accounts/%v", k.EndPoint.Authorization, username)).
		End()

	if errs != nil {
		k.Logger.Fatal("Could not retrieve new tokens")
	}

	switch resp.StatusCode {
	case 200:
		fmt.Println("Account information:")
		k.Logger.Info(fmt.Sprintf("Response Status: %v", resp.Status))
		k.Logger.Info(fmt.Sprintf("Response: %v", resp))
		k.Logger.Info(fmt.Sprintf("body: %v", body))
	default:
		k.Logger.Info(fmt.Sprintf("Response Status: %v", resp.Status))
		k.Logger.Info(fmt.Sprintf("Response: %v", resp))
		k.Logger.Info(fmt.Sprintf("body: %v", body))
	}
}
