package kth

import "sort"

func kthLargest(nums []int, k int) int {
	return useSort(nums, k)
}

// useSort time complexity O(nlgn), space complexity O(1)
func useSort(nums []int, k int) int {
	n := len(nums)
	if n < 1 || k > n {
		return 0
	}
	sort.Ints(nums)
	return nums[n-k]
}
