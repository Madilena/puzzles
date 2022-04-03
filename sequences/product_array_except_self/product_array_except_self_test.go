package main

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	test := []int{-1, 1, 0, -3, 3}
	expected := []int{0, 0, 9, 0, 0}
	actual := productExceptSelf(test)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
