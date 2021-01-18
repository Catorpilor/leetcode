package sum

import "sort"

func maxKSumOps(nums []int, k int) int {
	return useTwoPointers(nums, k)
}

// useTwoPointers time complexity O(N*log(N)), space complexity O(1)
func useTwoPointers(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	l, r := 0, n-1
	var ans int
	for l < r {
		cs := nums[l] + nums[r]
		if cs == k {
			ans++
			l++
			r--
		} else if cs < k {
			l++
		} else {
			r--
		}
	}
	return ans
}
