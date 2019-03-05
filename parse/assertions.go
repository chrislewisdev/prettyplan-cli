package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func expectString(t *testing.T, expected, actual string) {
	if actual != expected {
		t.Error("Expected '" + expected + "' but got: '" + actual + "'")
	}
}

func expectStrings(t *testing.T, expected, actual []string) {
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected ", expected, " but got: ", actual)
	}
}

func expectInt(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Error(fmt.Sprintf("Expected '%d' but got: '%d'", expected, actual))
	}
}

func expectBool(t *testing.T, expected, actual bool) {
	if actual != expected {
		t.Error(fmt.Sprintf("Expected '%v' but got: '%v'", expected, actual))
	}
}
