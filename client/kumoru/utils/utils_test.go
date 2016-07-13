package utils

import "testing"

func TestFormatTime(t *testing.T) {
	if FormatTime("2016-07-13T13:29:49.954048") != "Sun, 31 Dec 0000 18:00:00 CST" {
		t.Error("expected: Sun, 31 Dec 0000 18:00:00 CST")
	}
}
