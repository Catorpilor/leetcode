package lhs

import "github.com/catorpilor/leetcode/utils"

func findLHS(nums []int) int {
	return useHashmap(nums)
}

// useHashmap time complexity O(N), space compelxity O(N)
func useHashmap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	set := make(map[int]int, n)
	for i := range nums {
		set[nums[i]]++
	}
	var ans int
	for _, num := range nums {
		if v, exists := set[num+1]; exists {
			ans = utils.Max(ans, v+set[num])
		}
	}
	return ans
}
