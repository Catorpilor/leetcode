package move

import (
	"sort"
)

func MinMoves(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	t, c, e := 0, 0, len(nums)-1
	// nums[0] is the minValue
	// nums[e] is the max value
	// when nums[0] = nums[e] exit the for loop
	for nums[e] != nums[0] {
		c = nums[e] - nums[0]
		t += c
		nums[0] += c
		e -= 1
		nums[e] += t
	}
	return t
}

func MinMoves2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// math equation
	// suppose after m moves all elements in array equal to x
	// sum + m * (n-1) = x * n -- <1>
	// and suppose minNum is the minest in the nums
	// x in <1> equation actually equals to minNum + m
	// so finally we have
	// m = sum - minNum * n
	mv, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		mv = min(mv, nums[i])
	}
	return sum - mv*len(nums)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
