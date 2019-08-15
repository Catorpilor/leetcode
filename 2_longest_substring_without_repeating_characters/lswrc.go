package lswrc

import (
	"github.com/catorpilor/leetcode/utils"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	// sliding window
	left, right, max := 0, 0, 0
	hset := make(map[byte]bool)
	for right < n {
		if hset[s[right]] {
			// found in the set
			delete(hset, s[left])
			left++
		} else {
			hset[s[right]] = true
			right++
			max = utils.Max(max, right-left)
		}
	}
	return max
}

func bruteForceLength(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if allUnique(s, i, j) {
				max = utils.Max(max, j-i)
			}
		}
	}
	return max
}
func allUnique(s string, l, r int) bool {
	hset := make(map[byte]bool)
	for i := l; i < r; i++ {
		if hset[s[i]] {
			return false
		}
		hset[s[i]] = true
	}
	return true
}

func lengthOfLongestSubstringOpt(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	max := 0
	hset := make(map[byte]int)
	l, r := 0, 0
	for r < n {
		// expand [l,r]
		if k, ok := hset[s[r]]; ok {
			l = utils.Max(l, k)
		}
		max = utils.Max(max, r-l+1)
		hset[s[r]] = r + 1
		r++
	}
	return max
}
