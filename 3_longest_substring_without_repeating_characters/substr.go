package substr

import (
	"github.com/catorpilor/leetcode/utils"
)

func longestSubstr(s string) int {
	return useTwoPointers2(s)
}

// useTwoPointers time complexity O(N), space complexity O(N)
func useTwoPointers(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	var ans int
	l, r := 0, 1
	seen := make(map[byte]int, n)
	seen[s[l]] = l
	var prev int
	for r < n {
		if pos, exists := seen[s[r]]; exists {
			prev = l
			l = pos + 1
			if l == r {
				seen = make(map[byte]int, n)
			} else {
				for i := prev; i < l; i++ {
					delete(seen, s[i])
				}
			}
		}
		seen[s[r]] = r
		// fmt.Printf("l:%d, r:%d, s[r]=%s, ans: %d\n", l, r, string(s[r]), ans)
		ans = utils.Max(ans, r-l+1)
		r++
	}
	return ans
}

// useTwoPointers2 time complexity O(N), space complexity O(N)
func useTwoPointers2(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	var ans int
	var l, r int
	seen := make(map[byte]bool, n)
	// a sliding windonw technique
	for r < n {
		if seen[s[r]] {
			// shrink the window by moving l forward
			delete(seen, s[l])
			l++
		} else {
			seen[s[r]] = true
			ans = utils.Max(ans, r-l+1)
			r++
		}
	}
	return ans
}
