package kumoru

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEndpoints(t *testing.T) {
	os.Clearenv()

	p, err := LoadEndpoints("example-cfg.ini", "endpoints")

	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "pool.kumoru.io:5000", p.Pool, "Expect pool endpoint to match")
	assert.Equal(t, "application.kumoru.io", p.Application, "Expect application endpoint to match")
	assert.Equal(t, "authorization.kumoru.io:5000", p.Authorization, "Expect authorization endpoint to match")
}

func TestLoadEndpointsFileNotFound(t *testing.T) {
	os.Clearenv()
	p, err := LoadEndpoints("fake-file.ini", "")

	assert.NotNil(t, err, "Expecting an error")

	assert.Equal(t, "", p.Pool, "Expect pool endpoint to match")
	assert.Equal(t, "", p.Application, "Expect application endpoint to match")
	assert.Equal(t, "", p.Authorization, "Expect authorization endpoint to match")
}
