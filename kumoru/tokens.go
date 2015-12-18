package kumoru

import (
	"gopkg.in/ini.v1"
)

type Ktokens struct {
	Public  string
	Private string
}

func LoadTokens(filename string, section string) (Ktokens, error) {
	config, err := ini.Load(filename)
	if err != nil {
		return Ktokens{}, err
	}

	iniTokens, err := config.GetSection(section)
	if err != nil {
		return Ktokens{}, err
	}

	pubToken := iniTokens.Key("kumoru_token_public").String()
	pvToken := iniTokens.Key("kumoru_token_private").String()

	return Ktokens{
		Public:  pubToken,
		Private: pvToken,
	}, nil

}
