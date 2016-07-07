package location

import (
	"reflect"
	"testing"
)

func TestBuildFindPath(t *testing.T) {

	cases := []struct {
		location Location
		expected string
	}{
		{
			location: Location{},
			expected: "https://locations.kumoru.io/v1/locations/",
		},
		{
			location: Location{
				Provider: "amazon",
			},
			expected: "https://locations.kumoru.io/v1/locations/amazon",
		},
		{
			location: Location{
				Provider: "amazon",
				Region:   "us-east-1",
			},
			expected: "https://locations.kumoru.io/v1/locations/amazon/us-east-1",
		},
	}

	for _, c := range cases {
		result := c.location.buildFindPath("https://locations.kumoru.io")

		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("result == %v, expected %v", result, c.expected)
		}
	}
}
