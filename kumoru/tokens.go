package kumoru

import (
	"github.com/vaughan0/go-ini"
)

type Ktokens struct {
	Public  string
	Private string
}

func LoadTokens(filename string, section string) (Ktokens, error) {
	config, err := ini.LoadFile(filename)
	if err != nil {
		return Ktokens{}, err
	}

	iniTokens := config.Section(section)

	pubToken := iniTokens["kumoru_token_public"]
	pvToken := iniTokens["kumoru_token_private"]

	return Ktokens{
		Public:  pubToken,
		Private: pvToken,
	}, nil

}
