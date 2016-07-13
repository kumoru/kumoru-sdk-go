package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	kumoru := m.Run()
	os.Exit(kumoru)
}
