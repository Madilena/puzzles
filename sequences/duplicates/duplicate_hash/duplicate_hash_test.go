package main

import (
	"reflect"
	"testing"
)

func TestDuplicateHash(t *testing.T) {
	nums := []int{0, 6, 10, 25, 33, 25}
	expected := []int{25}
	actual := duplicateHash(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDuplicateZero(t *testing.T) {
	nums := []int{0, 0, 10, 25, 33, 25}
	expected := []int{0, 25}
	actual := duplicateHash(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDuplicateNegative(t *testing.T) {
	var nums = []int{0, -1, -1, 10, 33, 25}
	var expected = []int{-1}
	actual := duplicateHash(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDuplicateEmpty(t *testing.T) {
	var nums []int
	var expected []int
	actual := duplicateHash(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestDuplicateNone(t *testing.T) {
	var nums = []int{0, -1, 10, 33, 25}
	var expected []int
	actual := duplicateHash(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}