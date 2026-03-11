package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// Convert to lowercase and remove spaces for a better check
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	length := len(s)
	for i := 0; i < length/2; i++ {
		// Compare characters from both ends
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	word := "Madam"
	if isPalindrome(word) {
		fmt.Printf("'%s' is a palindrome!\n", word)
	} else {
		fmt.Printf("'%s' is NOT a palindrome.\n", word)
	}
}
