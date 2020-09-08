package common

import (
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
	result := Getenv("USER", "false")
	expected := os.Getenv("USER")
	if result != expected {
		t.Errorf("got %s expected %s", result, expected)
	}
}
