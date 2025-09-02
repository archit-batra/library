package main

import (
	"unicode"
)

// isPalindrome checks if s is a palindrome ignoring non-alphanum and case.
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		// Skip non-alphanumeric from left
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
		}
		// Skip non-alphanumeric from right
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
		}
		if left < right {
			if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
				return false // Mismatch
			}
			left++
			right--
		}
	}
	return true
}
