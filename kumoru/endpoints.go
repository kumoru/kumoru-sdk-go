package kumoru

import (
	"github.com/vaughan0/go-ini"
)

type Endpoints struct {
	Application   string
	Authorization string
	Pool          string
}

func LoadEndpoints(filename string, section string) (Endpoints, error) {
	config, err := ini.LoadFile(filename)
	if err != nil {
		return Endpoints{}, err
	}

	iniEndpoints := config.Section(section)

	application_api := iniEndpoints["kumoru_application_api"]
	authorization_api := iniEndpoints["kumoru_authorization_api"]
	pool_api := iniEndpoints["kumoru_pool_api"]

	return Endpoints{
		Application:   application_api,
		Authorization: authorization_api,
		Pool:          pool_api,
	}, nil

}
