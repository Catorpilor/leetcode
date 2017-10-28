package robber

import "github.com/catorpilor/leetcode/utils"

func Robber(nums []int) int {
	// brute force
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	return utils.Max(rob(nums, 0, n-2), rob(nums, 1, n-1))

}

func rob(nums []int, l, r int) int {
	var include, exclude int
	for i := l; i <= r; i++ {
		t, e := include, exclude
		include = e + nums[i]
		exclude = utils.Max(t, e)
	}
	return utils.Max(include, exclude)
}
