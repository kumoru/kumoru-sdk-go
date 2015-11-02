package kumoru

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadCreds(t *testing.T) {
	os.Clearenv()

	p, err := LoadCreds("example-cfg.ini", "auth")

	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "USER", p.UserName, "Expect username to match")
	assert.Equal(t, "SECRET", p.Password, "Expect password to match")
}

func TestLoadCredsMissingUserName(t *testing.T) {
	os.Clearenv()

	p, err := LoadCreds("example-cfg.ini", "missing-username")

	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "", p.UserName, "Expect username Empty string match")
	assert.Equal(t, "", p.Password, "Expect password Empty string match")
}

func TestLoadCredsMissingPassword(t *testing.T) {
	os.Clearenv()

	p, err := LoadCreds("example-cfg.ini", "missing-password")

	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "", p.UserName, "Expect username Empty string match")
	assert.Equal(t, "", p.Password, "Expect password Empty string match")
}

func TestLoadCredsFileNotFound(t *testing.T) {
	os.Clearenv()
	p, err := LoadCreds("fake-file.ini", "")

	assert.NotNil(t, err, "Expecting an error")

	assert.Equal(t, "", p.UserName, "Expect username Empty string match")
	assert.Equal(t, "", p.Password, "Expect password Empty string match")
}
