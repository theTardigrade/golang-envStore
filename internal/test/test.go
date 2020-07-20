package test

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, name string, expectedValue, actualValue interface{}) {
	if reflect.DeepEqual(expectedValue, actualValue) {
		return
	}

	t.Errorf("Expected %s to be %#v, but got %#v.", name, expectedValue, actualValue)
}
