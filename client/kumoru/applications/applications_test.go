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

package applications

import (
	"reflect"
	"testing"
)

func TestTransformRules(t *testing.T) {
	cases := []struct {
		rules       []string
		expected    map[string]int
		expectedErr error
	}{
		{
			rules:       []string{"latest=100"},
			expected:    map[string]int{"latest": 100},
			expectedErr: nil,
		},
		{
			rules:       []string{"latest=50", "alpha=50"},
			expected:    map[string]int{"latest": 50, "alpha": 50},
			expectedErr: nil,
		},
	}

	for _, c := range cases {
		result := transformRules(&c.rules)

		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("result == %v, expected %v", result, c.expected)
		}
	}

}

func TestEnvironment(t *testing.T) {
	cases := []struct {
		envFile     string
		enVars      []string
		expected    map[string]string
		expectedErr error
	}{
		{
			envFile:     "",
			enVars:      []string{"FOO=bar", "DB=password=withequals"},
			expected:    map[string]string{"FOO": "bar", "DB": "password=withequals"},
			expectedErr: nil,
		},
	}
	for _, c := range cases {
		result := transformEnvironment(&c.envFile, &c.enVars)

		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("result == %v, expected %v", result, c.expected)
		}
	}
}

/*func TestReadCertificatesEmpty(t *testing.T) {

	var cert, key, ca string
	expected := ""

	result := readCertificates(&cert, &key, &ca)

	if result != expected {
		t.Errorf("result == %v, want %v", result, expected)
	}
}*/
//TODO Implement following test case with file reads
/*func TestReadCertificatesExists(t *testing.T) {

	cert := "mycert"
	key := "mykey"
	ca := "myca"

	expected := "{\"certificate\": \"mycert\", \"private_key\": \"mykey\", \"certificate_authority\": \"myca\"}"

	result := readCertificates(&cert, &key, &ca)

	if result != expected {
		t.Errorf("result == %v, want %v", result, expected)
	}
}*/
