package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "anagram"
	t := "nagaram"

	result := isAnagram(s, t)
	fmt.Println(result)
}

func isAnagram(s string, t string) bool {

	if !(len(s) == len(t)) {
		return false
	}

	if s == t {
		return true
	}
	s_sort := SortString(s)
	t_sort := SortString(t)

	if !(t_sort == s_sort) {
		return false
	} else {
		return true
	}
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
