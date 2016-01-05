package kumoru

import "gopkg.in/ini.v1"

type Endpoints struct {
	Application   string
	Authorization string
	Pool          string
}

func LoadEndpoints(filename string, section string) Endpoints {
	config, err := ini.Load(filename)
	if err != nil {
		return Endpoints{
			Application:   "https://application.kumoru.io",
			Authorization: "https://authorization.kumoru.io",
			Pool:          "https://pool.kumoru.io",
		}
	}

	iniEndpoints, err := config.GetSection(section)
	if err != nil {
		return Endpoints{
			Application:   "https://application.kumoru.io",
			Authorization: "https://authorization.kumoru.io",
			Pool:          "https://pool.kumoru.io",
		}
	}

	return Endpoints{
		Application:   iniEndpoints.Key("kumoru_application_api").String(),
		Authorization: iniEndpoints.Key("kumoru_authorization_api").String(),
		Pool:          iniEndpoints.Key("kumoru_pool_api").String(),
	}
}
