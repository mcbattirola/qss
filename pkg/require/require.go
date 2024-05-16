package require

import (
	"reflect"
	"testing"
)

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Expected no error but got error '%s'", err.Error())
	}
}

func Equal(t *testing.T, expected any, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v to be equal to %v, but result is different", expected, actual)
	}
}
