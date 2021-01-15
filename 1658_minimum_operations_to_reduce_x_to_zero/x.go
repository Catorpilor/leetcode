package x

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func minRes(nums []int, x int) int {
	// return useBruteForce(nums, x)
	return usePrefixSum(nums, x)
}

func useBruteForce(nums []int, x int) int {
	ans := helper(nums, x, len(nums))
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func helper(nums []int, x, ol int) int {
	n := len(nums)
	if x == 0 {
		return ol - n
	}
	if x < 0 || n == 0 {
		return math.MaxInt32
	}
	return utils.Min(helper(nums[1:], x-nums[0], ol), helper(nums[:n-1], x-nums[n-1], ol))
}

// usePrefixSum time complexity O(N),space complexity O(N)
func usePrefixSum(nums []int, x int) int {
	psum := make([]int, len(nums))
	copy(psum, nums)
	n := len(nums)
	for i := 1; i < n; i++ {
		psum[i] += psum[i-1]
	}
	target := psum[n-1] - x
	if target == 0 {
		return n
	}
	set := make(map[int]int, n)
	set[0] = -1
	var sum int
	res := math.MinInt32
	for i := 0; i < n; i++ {
		sum += nums[i]
		if p, exists := set[sum-target]; exists {
			res = utils.Max(res, i-p)
		}
		set[sum] = i
	}
	if res == math.MinInt32 {
		return -1
	}
	return n - res
}
