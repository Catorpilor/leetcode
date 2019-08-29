package lswt

import "github.com/catorpilor/LeetCode/utils"

// LengthOfLongestSubstring returns the longest substring with at most two distinct characters
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 2 {
		return n
	}
	// use slide window algorithm
	// here k=2
	beg, end, count, res := 0, 0, 0, 0
	cache := make(map[byte]int, n)
	for end < n {
		cache[s[end]]++
		if cache[s[end]] == 1 {
			// new character
			count++
		}
		end++
		for count > 2 {
			// more than two distinct characters
			// more beg from left to right
			cache[s[beg]]--
			if cache[s[beg]] == 0 {
				// remove one character
				count--
			}
			beg++
		}
		res = utils.Max(res, end-beg)

	}
	return res
}
