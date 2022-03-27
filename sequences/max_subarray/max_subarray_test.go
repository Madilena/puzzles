package main

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6
	actual := maxSubArray(arr)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
