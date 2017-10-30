package ts

import (
	"math"
)

func TargetSum(nums []int, s int) int {
	// brute force
	return helper(nums, 0, 0, s)
}

func helper(nums []int, i, sum, s int) int {
	if i == len(nums) {
		if sum == s {
			return 1
		} else {
			return 0
		}
	}
	var count int
	count += helper(nums, i+1, sum+nums[i], s)
	count += helper(nums, i+1, sum-nums[i], s)
	return count
}

func TargetSumWithMemo(nums []int, s int) int {
	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 2001)
		for j := range memo[i] {
			memo[i][j] = math.MinInt32
		}
	}
	return helperWithMemo(nums, 0, 0, s, memo)

}

func helperWithMemo(nums []int, i, sum, s int, memo [][]int) int {
	if i == len(nums) {
		if sum == s {
			return 1
		}
		return 0
	}
	if memo[i][sum+1000] != math.MinInt32 {
		return memo[i][sum+1000]
	}
	var count int
	count += helperWithMemo(nums, i+1, sum+nums[i], s, memo)
	count += helperWithMemo(nums, i+1, sum-nums[i], s, memo)
	memo[i][sum+1000] = count
	return count
}

func TargetSum2(nums []int, s int) int {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	if s > sum || (sum-s)%2 == 1 {
		return 0
	}
	target := (sum - s) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}
