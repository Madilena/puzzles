package main

import (
	"reflect"
	"testing"
)

func TestInValidParenthes(t *testing.T) {
	s := "([)]"
	expected := false
	actual := isValid(s)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %t, got %t", expected, actual)
	}
}

func TestValidParenthes(t *testing.T) {
	s := "{{}}()[()]"
	expected := true
	actual := isValid(s)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %t, got %t", expected, actual)
	}
}
