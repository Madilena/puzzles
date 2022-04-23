package main

import (
	"testing"
)

func TestAnagramHappy(t *testing.T) {
	s := "anagram"
	b := "nagaram"
	expected := true
	actual := isAnagram(s, b)

	if !(actual == expected) {
		t.Fatalf("Expected %t, got %t", expected, actual)
	}
}

func TestAnagramSad(t *testing.T) {
	s := "anagram"
	b := "candygram"
	expected := false
	actual := isAnagram(s, b)

	if !(actual == expected) {
		t.Fatalf("Expected %t, got %t", expected, actual)
	}
}