package kumoru

import (
	"gopkg.in/ini.v1"
)

type Endpoints struct {
	Application   string
	Authorization string
	Pool          string
}

func LoadEndpoints(filename string, section string) (Endpoints, error) {
	config, err := ini.Load(filename)
	if err != nil {
		return Endpoints{}, err
	}

	iniEndpoints, err := config.GetSection(section)
	if err != nil {
		return Endpoints{
			Application:   "https://application.kumoru.io:5000",
			Authorization: "https://authorization.kumoru.io:5000",
			Pool:          "https://pool.kumoru.io:5000",
		}, nil
	}

	return Endpoints{
		Application:   iniEndpoints.Key("kumoru_application_api").String(),
		Authorization: iniEndpoints.Key("kumoru_authorization_api").String(),
		Pool:          iniEndpoints.Key("kumoru_pool_api").String(),
	}, nil

}
