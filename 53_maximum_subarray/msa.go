package msa

import (
	"github.com/catorpilor/leetcode/utils"
)

func MaxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	sum, maxSofar := 0, nums[0]
	for _, c := range nums {
		sum += c
		maxSofar = utils.Max(sum, maxSofar)
		sum = utils.Max(sum, 0)
	}
	return maxSofar
}

func MaxSubArray2(nums []int) int {
	// dp
	// dp[j] represents the max sum so far end with  jth
	// element in nums
	// so for dp[i] = dp[i-1] > 0 : dp[i-1] : 0 + nums[i]
	// O(n) space and O(n) time
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSofar := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i]
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
		maxSofar = utils.Max(maxSofar, dp[i])
	}
	return maxSofar
}
