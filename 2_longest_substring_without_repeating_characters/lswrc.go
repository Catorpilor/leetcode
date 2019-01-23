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
