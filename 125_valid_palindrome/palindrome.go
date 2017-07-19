package palindrome

import (
	"strings"
)

func isValidCharacter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= '0' && c <= '9'
}

func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !isValidCharacter(s[i]) {
			i++
			continue
		}
		if !isValidCharacter(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		} else {
			i, j = i+1, j-1
		}
	}
	return true
}
