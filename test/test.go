package test

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, name string, expectedValue, actualValue interface{}) {
	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Errorf("Expected %s to be \"%s\", but got \"%s\".", name, expectedValue, actualValue)
	}
}
