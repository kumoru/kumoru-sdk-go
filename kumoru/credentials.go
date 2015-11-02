package kumoru

import (
	"github.com/vaughan0/go-ini"
)

type Creds struct {
	Password string
	UserName string
}

func LoadCreds(filename string, section string) (Creds, error) {
	config, err := ini.LoadFile(filename)
	if err != nil {
		return Creds{}, err
	}

	iniAuth := config.Section(section)

	username, ok := iniAuth["kumoru_username"]
	if !ok {
		return Creds{}, err
	}

	password, ok := iniAuth["kumoru_password"]
	if !ok {
		return Creds{}, err
	}

	return Creds{
		Password: password,
		UserName: username,
	}, nil

}
