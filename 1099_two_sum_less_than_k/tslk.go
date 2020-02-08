package tslk

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func twoSumLessThanK(nums []int, k int) int {
	res := -1
	n := len(nums)
	if n < 2 {
		return res
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	l, r := 0, n-1
	for l < r {
		if nums[l]+nums[r] < k {
			res = utils.Max(res, nums[l]+nums[r])
			l++
		} else {
			r--
		}
	}
	return res
}
