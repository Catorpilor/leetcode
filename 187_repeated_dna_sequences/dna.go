package dna

import (
	"strings"
)

func findRepeatedSeq(s string) []string {
	return useBruteForce(s)
}

// useBruteForce time complexity O(N^2), space complexity O(1)
func useBruteForce(s string) []string {
	n := len(s)
	if n <= 10 {
		return nil
	}
	var ans []string
	found := make(map[string]bool, n)
	for i := 0; i < n-9; i++ {
		tmp := s[i : i+10]
		if !found[tmp] && strings.Index(s[i+1:], tmp) != -1 {
			found[tmp] = true
			ans = append(ans, tmp)
		}

	}
	return ans
}
