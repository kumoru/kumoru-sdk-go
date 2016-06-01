/*
Copyright 2016 Kumoru.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
