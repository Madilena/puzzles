package main

import (
	"reflect"
	"testing"
)

func TestPalindromeHappy(t *testing.T) {
	s := "racecar"
	expected := true
	actual := isPalindrome(s)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %t, got %t", expected, actual)
	}
}

func TestPalindromeFunny(t *testing.T) {
	s := "go hang a salami im a lasagna hog"
	expected := true
	actual := isPalindrome(s)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %t, got %t", expected, actual)
	}
}

func TestPalindromeSad(t *testing.T) {
	s := "too"
	expected := false
	actual := isPalindrome(s)

	if !(reflect.DeepEqual(actual, expected)) {
		t.Fatalf("Expected %t, got %t", expected, actual)
	}
}