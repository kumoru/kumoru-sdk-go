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
