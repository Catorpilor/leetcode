package tss

import (
	"sort"
)

func sumSmaller(nums []int, target int) int {
	return bruteForce(nums, target)
}

// bruteForce time complexity O(n^3) space complexity O(1)
func bruteForce(nums []int, target int) int {
	var res int
	n := len(nums)
	if n < 3 {
		return res
	}
	// sort nums
	// [-2,0,1,3]
	sort.Slice(nums, func(i, j int) bool { return nums[i] <= nums[j] })
	for i := n - 1; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				if nums[i]+nums[j]+nums[k] < target {
					res += k + 1
					break
				}
			}
		}
	}
	return res
}
