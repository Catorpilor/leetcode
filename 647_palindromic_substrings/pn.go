package pn

import "fmt"

func CountSubStrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var count int
	for i := 0; i < len(s); i++ {
		checkPanlidrome(s, i, i, &count)
		checkPanlidrome(s, i, i+1, &count)
	}
	return count
}

func checkPanlidrome(s string, l, r int, count *int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		*count, l, r = *count+1, l-1, r+1
	}
}

// useBruteForce time complexity O(N^2), space complexity O(1)
func useBruteForce(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	ans := n
	for k := 2; k <= n; k++ {
		for i := 0; i+k <= n; i++ {
			if isPalindrome(s[i : i+k]) {
				fmt.Println(s[i : i+k])
				ans++
			}
		}
	}
	return ans
}

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
