package subSumk

import "github.com/catorpilor/LeetCode/utils"

func maxSubLen(nums []int, k int) int {
	return preSumWithHash(nums, k)
}

// preSumWithHash time complexity O(2n) space complexity O(n)
func preSumWithHash(nums []int, k int) int {
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
	// key is the sum and value is the index
	set := make(map[int]int, n)
	// edge case,
	// when sum is 0, the length of sub array is index - (-1)
	set[0] = -1
	for i := range preSum {
		if v, exists := set[preSum[i]-k]; exists {
			res = utils.Max(res, i-v)
		}
		if _, exists := set[preSum[i]]; !exists {
			set[preSum[i]] = i
		}
	}
	return res
}

// bruteForce time complexity is O(n^2) space complexity is O(1)
func bruteForce(nums []int, k int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	var res, sum int
	for i := range nums {
		sum = nums[i]
		if sum == k {
			res = utils.Max(res, 1)
		}
		for j := i + 1; j < n; j++ {
			sum += nums[j]
			if sum == k {
				res = utils.Max(res, j-i+1)
			}
		}
	}
	return res
}
