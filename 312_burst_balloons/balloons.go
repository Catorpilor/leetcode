package balloon

import (
	"github.com/catorpilor/leetcode/utils"
)

func MaxCoins(nums []int) int {
	n := len(nums)
	// memorization dfs
	// dp[i][j] represents the max coins from balloons i to j
	dp := [501][501]int{}
	return dfs(nums, &dp, 0, n-1, 1, 1)
}

func dfs(nums []int, ans *[501][501]int, i, j, left, right int) int {
	if i > j {
		return 0
	}
	if (*ans)[i][j] != 0 {
		return (*ans)[i][j]
	}
	res := 0
	for k := i; k <= j; k++ {
		// if k+1 == len(nums) {
		// 	curRight = 1
		// } else {
		// 	curRight = nums[k+1]
		// }
		res = utils.Max(res, left*nums[k]*right+dfs(nums, ans, k+1, j, nums[k], right)+dfs(nums, ans, i, k-1, left, nums[k]))
	}
	(*ans)[i][j] = res
	return res
}
