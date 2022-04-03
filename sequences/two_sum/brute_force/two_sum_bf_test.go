package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{0, 6, 10, 25}
	target := 31
	expected := [][]int{[]int{1, 3}}
	actual := twoSumBf(nums, target)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestTwoSumMiddle(t *testing.T) {
	nums := []int{0, 6, 10, 25}
	target := 16
	expected := [][]int{[]int{1, 2}}
	actual := twoSumBf(nums, target)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestTwoSumBothBoundary(t *testing.T) {
	nums := []int{0, 6, 10, 11, 50, 25}
	target := 25
	expected := [][]int{[]int{0, 5}}
	actual := twoSumBf(nums, target)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestTwoSumZero(t *testing.T) {
	nums := []int{6, 0, 10, 0, 50, 25}
	target := 0
	expected := [][]int{[]int{1, 3}}
	actual := twoSumBf(nums, target)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestTwoResultsZero(t *testing.T) {
	nums := []int{6, 0, 10, 25, 50, 25}
	target := 50
	expected := [][]int{[]int{1, 4}, []int{3, 5}}
	actual := twoSumBf(nums, target)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestThreeResultsZero(t *testing.T) {
	nums := []int{6, 1, 7, 4, 2, 10, 5}
	target := 11
	expected := [][]int{[]int{0, 6}, []int{1, 5}, []int{2, 3}}
	actual := twoSumBf(nums, target)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
