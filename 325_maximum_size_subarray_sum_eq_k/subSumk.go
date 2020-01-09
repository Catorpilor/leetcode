package subSumk

import "github.com/catorpilor/LeetCode/utils"

func maxSubLen(nums []int, k int) int {
	return bruteForce(nums, k)
}

func bruteForce(nums []int, k int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	preSum := make([]int, n)
	preSum[0] = nums[0]
	for i := 1; i < n; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	var res int
	return res
}
