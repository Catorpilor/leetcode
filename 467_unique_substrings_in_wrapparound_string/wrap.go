package wrap

import (
	"github.com/catorpilor/leetcode/utils"
)

func NumOfSubStrings(p string) int {
	n := len(p)
	if n <= 1 {
		return n
	}
	// dp[i] represents the max unique substring ending with ith
	// 0 - 'a', ....25 - 'z'
	dp := make([]int, 26)
	maxLen := 0
	for i := range p {
		if i > 0 && (p[i]-p[i-1] == 1 || p[i-1]-p[i] == 25) {
			maxLen += 1
		} else {
			maxLen = 1
		}
		idx := p[i] - 'a'
		// just store the longest substring
		// for example abcdbcd
		// the longest substring ending with 'd' is 4 'abcd' which overlap the 2nd 'bcd'
		dp[idx] = utils.Max(dp[idx], maxLen)
	}
	var ret int
	for i := range dp {
		ret += dp[i]
	}
	return ret
}

func NumOfSubStrings2(p string) int {
	// brute force
	n := len(p)
	if n <= 1 {
		return n
	}
	hash := make(map[string]bool)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if j > i && (p[j]-p[j-1] != 1 && p[j-1]-p[j] != 25) {
				break
			}
			if j+1 < n {
				hash[p[i:j+1]] = true
			} else {
				hash[p[i:]] = true
			}
		}
	}
	return len(hash)
}
