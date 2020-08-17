package eat

import (
	"github.com/catorpilor/leetcode/utils"
)

func minDays(n int) int {
	return useHashmap(n)
}

// useHashmap time complexity O(lgN), space complexity O(N)
func useHashmap(n int) int {
	set := make(map[int]int, n)
	return helper(set, n)
}

func helper(set map[int]int, n int) int {
	if n <= 1 {
		return n
	}
	if _, exists := set[n]; !exists {
		set[n] = 1 + utils.Min(n%2+helper(set, n/2), n%3+helper(set, n/3))
	}
	return set[n]
}
