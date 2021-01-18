package sum

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

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

// useHashmap time complexity O(N), space complexity O(N)
func useHashmap(nums []int, k int) int {
	n := len(nums)
	set := make(map[int]int, n)
	for _, num := range nums {
		set[num]++
	}
	var ans int
	for num, v := range set {
		if k-num == num {
			ans += v / 2
		} else if set[k-num] != 0 {
			cm := utils.Min(v, set[k-num])
			ans += cm
			set[num] -= cm
			set[k-num] -= cm
		}
	}
	return ans
}
