package main

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	actual := threeSum(nums)

	if !(actual == expected) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
