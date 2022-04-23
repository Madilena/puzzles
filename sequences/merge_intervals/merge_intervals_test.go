package main

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	actual := merge(nums)

	if !(actual == expected) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestMergeIntervalsAgain(t *testing.T) {
	nums := [][]int{{1, 4}, {4, 5}}
	expected := [][]int{{1, 5}}
	actual := merge(nums)

	if !(actual == expected) {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
