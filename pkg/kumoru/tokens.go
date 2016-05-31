package kumoru

import (
	"os"

	"gopkg.in/ini.v1"
)

// Ktokens contains public and private tokens
type Ktokens struct {
	Public  string
	Private string
}

// LoadTokens from a file returning a struct of type Ktokens
func LoadTokens(filename string, section string) (Ktokens, error) {
	config, err := ini.Load(filename)
	if err != nil {
		return Ktokens{}, err
	}

	iniTokens, err := config.GetSection(section)
	if err != nil {
		return Ktokens{}, err
	}

	return Ktokens{
		Public:  iniTokens.Key("kumoru_token_public").String(),
		Private: iniTokens.Key("kumoru_token_private").String(),
	}, nil

}

// SaveTokens writes tokens to a file
func SaveTokens(directory, filename, section string, tokens Ktokens) error {
	kfile := directory + filename

	config, err := ini.Load(kfile)

	if err != nil {
		os.Mkdir(directory, 0755)
		config = ini.Empty()
		config.NewSection(section)
	}

	config.Section(section).NewKey("kumoru_token_public", tokens.Public)
	config.Section(section).NewKey("kumoru_token_private", tokens.Private)
	return config.SaveTo(kfile)
}

// HasTokens checks a file to make sure there are tokens stored
func HasTokens(filename, section string) bool {
	config, err := ini.Load(filename)
	if err != nil {
		return false
	}

	tokens, err := config.GetSection(section)

	if err != nil {
		return false
	}

	if tokens.Key("kumoru_token_public").String() == "" || tokens.Key("kumoru_token_private").String() == "" {
		return false
	}

	return true

}
