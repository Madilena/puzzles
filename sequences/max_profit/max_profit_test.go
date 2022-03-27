package main

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	nums := []int{1, 6, 4, 3, 7}
	expected := 6
	actual := maxProfit(nums)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestMaxProfitZero(t *testing.T) {
	nums := []int{7, 6, 4, 3, 1}
	expected := 0
	actual := maxProfit(nums)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestMaxProfitBoundaries(t *testing.T) {
	nums := []int{1, 2}
	expected := 1
	actual := maxProfit(nums)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestMaxProfitBoundariesA(t *testing.T) {
	nums := []int{1, 4, 2}
	expected := 3
	actual := maxProfit(nums)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestMaxProfitBoundariesMix(t *testing.T) {
	nums := []int{2, 1, 4}
	expected := 3
	actual := maxProfit(nums)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestMaxProfitDuplicate(t *testing.T) {
	nums := []int{2, 1, 2, 1, 0, 1, 2}
	expected := 2
	actual := maxProfit(nums)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestMaxProfitResetMax(t *testing.T) {
	nums := []int{3, 3, 5, 0, 0, 3, 1, 4}
	expected := 4
	actual := maxProfit(nums)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
func TestMaxProfitResetResult(t *testing.T) {
	nums := []int{3, 2, 6, 5, 0, 3}
	expected := 4
	actual := maxProfit(nums)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
