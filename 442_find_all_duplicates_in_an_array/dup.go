package dup

import "github.com/catorpilor/leetcode/utils"

func findDuplicates(nums []int) []int {
	// O(n) space
	n := len(nums)
	ret := make([]int, 0)
	if n <= 1 {
		return ret
	}
	hash := make(map[int]int)
	for i := range nums {
		hash[nums[i]] += 1
		if hash[nums[i]] > 1 {
			ret = append(ret, nums[i])
		}
	}
	return ret
}

// useConstantSpace time complexity O(N), space complexity O(1)
func useConstantSpace(nums []int) []int {
	// no more extra space
	// since  1<=nums[i]<=n
	// we can flip the num at pos nums[i] - 1
	// then if the number at pos nums[i] - 1 is negative, nums[i] is the number
	// occurs twice
	ret := make([]int, 0)
	for i := range nums {
		idx := utils.Abs(nums[i]) - 1
		if nums[idx] < 0 {
			ret = append(ret, idx+1)
		}
		nums[idx] = -nums[idx]
	}
	return ret

}
