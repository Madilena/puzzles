package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "racecar"
	result := isPalindrome(s)
	fmt.Println(result)
}

func isPalindrome(text string) bool {
	s := strings.ReplaceAll(text, " ", "")
	for i := 0; i < len(s); i++ {
		forward := string(s[i])
		backward := string(s[len(s)-1-i])
		if !(forward == backward) {
			return false
		}
	}
	return true
}
