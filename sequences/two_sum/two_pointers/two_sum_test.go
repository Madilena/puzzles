package main

import (
	"testing"
)

func TestTwoSumFast(t *testing.T) {
	nums := []int{0, 6, 10, 25}
	target := 31
	expected := []int{1, 3}
	actual := twoSum(nums, target)

	if !(actual == expected) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}