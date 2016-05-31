package kumoru

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEndpoints(t *testing.T) {
	os.Clearenv()

	p := LoadEndpoints("example-cfg.ini", "endpoints")

	assert.Equal(t, "http://pool.kumoru.io:5000", p.Pool, "Expect pool endpoint to match")
	assert.Equal(t, "http://application.kumoru.io", p.Application, "Expect application endpoint to match")
	assert.Equal(t, "http://authorization.kumoru.io:5000", p.Authorization, "Expect authorization endpoint to match")
}

func TestLoadEndpointsFileNotFound(t *testing.T) {
	os.Clearenv()
	p := LoadEndpoints("fake-file.ini", "")

	assert.NotNil(t, p, "Expecting an error")

	assert.Equal(t, "https://pool.kumoru.io", p.Pool, "Expect pool endpoint to match")
	assert.Equal(t, "https://application.kumoru.io", p.Application, "Expect application endpoint to match")
	assert.Equal(t, "https://authorization.kumoru.io", p.Authorization, "Expect authorization endpoint to match")
}
