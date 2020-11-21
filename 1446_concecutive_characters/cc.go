package cc

import "github.com/catorpilor/leetcode/utils"

func maxPower(s string) int {
	return useOnepass(s)
}

// useOnepass time complexity O(N), space complexity O(1)
func useOnepass(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	ans, cur := 1, 1
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			cur = 1
			continue
		}
		cur++
		ans = utils.Max(ans, cur)
	}
	return ans
}
