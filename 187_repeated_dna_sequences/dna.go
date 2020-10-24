package dna

import (
	"strings"
)

func findRepeatedSeq(s string) []string {
	// return useBruteForce(s)
	return useBits(s)
}

// useBruteForce time complexity O(N^2), space complexity O(n)
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

// useBits time complexity O(N), space complexity O(N)
func useBits(s string) []string {
	n := len(s)
	if n <= 10 {
		return nil
	}
	var ans []string
	occ := make(map[int]int)
	for i, t := 0, 0; i < n; i++ {
		t = t<<3&0x3FFFFFFF | int(s[i]&7)
		if occ[t] == 1 {
			ans = append(ans, s[i-9:i+1])
		}
		occ[t]++
	}
	return ans
}
