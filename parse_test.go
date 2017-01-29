package env

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	key, value, err := parseLine("x=y")

	if err != nil {
		t.Error(err)
	}

	if key != "X" {
		t.Error("Expected key to be \"X\", got \"" + key + "\".")
	}

	if value != "y" {
		t.Error("Expected key to be \"y\", got \"" + value + "\".")
	}
}
