package lswk

import "github.com/catorpilor/LeetCode/utils"

// LongestSubstringAtmostK returns the longest substring with at most K distinct characters
func LongestSubstringAtmostK(s string, k int) int {
	if k < 1 {
		return 0
	}
	n := len(s)
	if k >= n {
		return n
	}
	beg, end, count, res := 0, 0, 0, 0
	cache := make(map[byte]int, n)
	for end < n {
		cache[s[end]]++
		if cache[s[end]] == 1 {
			// new character
			count++
		}
		end++
		for count > k {
			cache[s[beg]]--
			if cache[s[beg]] == 0 {
				count--
			}
			beg++
		}
		res = utils.Max(res, end-beg)
	}
	return res
}
