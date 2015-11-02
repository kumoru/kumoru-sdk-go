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
