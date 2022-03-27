package main

import (
	"testing"
)

func TestContainsNoDuplicate(t *testing.T) {
	nums := []int{0, 6, 10, 25}
	expected := false
	actual := containsDuplicate(nums)

	if !(expected == actual) {
		t.Fatalf("Expected %t, got %t", expected, actual)
	}
}

func TestContainsDuplicate(t *testing.T) {
	nums := []int{0, 6, 10, 25, 0}
	expected := true
	actual := containsDuplicate(nums)

	if !(expected == actual) {
		t.Fatalf("Expected %t, got %t", expected, actual)
	}
}
