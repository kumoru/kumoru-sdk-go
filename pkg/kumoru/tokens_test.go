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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadKtokens(t *testing.T) {
	os.Clearenv()

	f, err := LoadTokens("example-cfg.ini", "tokens")

	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "PUBLIC_TOKEN", f.Public, "Expect Public Token to match")
	assert.Equal(t, "PRIVATE_TOKEN", f.Private, "Expect Private Token to match")
}

func TestLoadTokensFileNotFound(t *testing.T) {
	os.Clearenv()
	f, err := LoadTokens("fake-file.ini", "")

	assert.NotNil(t, err, "Expecting an error")

	assert.Equal(t, "", f.Public, "Expect Public Token to match")
	assert.Equal(t, "", f.Private, "Expect Private Token to match")
}
