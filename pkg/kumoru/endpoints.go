package kumoru

import (
	"os"

	"gopkg.in/ini.v1"
)

// Endpoints struct for all api services
type Endpoints struct {
	Application   string
	Authorization string
	Pool          string
}

// LoadEndpoints returns and Endpoints struct by reading them from a file or from defaults
func LoadEndpoints(filename string, section string) Endpoints {
	config, err := ini.Load(filename)

	appManagerURL := "https://application.kumoru.io"
	if os.Getenv("APPLICATION_MANAGER_URL") != "" {
		appManagerURL = os.Getenv("APPLICATION_MANAGER_URL")
	}

	authManagerURL := "https://authorization.kumoru.io"
	if os.Getenv("AUTHORIZATION_MANAGER_URL") != "" {
		authManagerURL = os.Getenv("AUTHORIZATION_MANAGER_URL")
	}

	poolManagerURL := "https://pool.kumoru.io"
	if os.Getenv("POOL_MANAGER_URL") != "" {
		poolManagerURL = os.Getenv("POOL_MANAGER_URL")
	}

	if err != nil {
		return Endpoints{
			Application:   appManagerURL,
			Authorization: authManagerURL,
			Pool:          poolManagerURL,
		}
	}

	iniEndpoints, err := config.GetSection(section)
	if err != nil {
		return Endpoints{
			Application:   appManagerURL,
			Authorization: authManagerURL,
			Pool:          poolManagerURL,
		}
	}

	return Endpoints{
		Application:   iniEndpoints.Key("kumoru_application_api").String(),
		Authorization: iniEndpoints.Key("kumoru_authorization_api").String(),
		Pool:          iniEndpoints.Key("kumoru_pool_api").String(),
	}
}
