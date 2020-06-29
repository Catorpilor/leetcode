package sum

import "sort"

func numSubSeq(nums []int, target int) int {
	return useTwoPointers(nums, target)
}

// useTwoPointers time complexity O(N*Log(N)), space complexity O(N)
func useTwoPointers(nums []int, target int) int {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	mod := 1e9 + 7
	power := make([]int, n)
	power[0] = 1
	for i := 1; i < n; i++ {
		power[i] = (power[i-1] * 2) % int(mod)
	}
	l, r := 0, n-1
	var ans int
	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			ans = (ans + power[r-l]) % int(mod)
			l++
		}
	}
	return ans
}
