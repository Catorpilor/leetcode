package robber

import (
	"github.com/catorpilor/leetcode/utils"
)

func Rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[l-1]
	}
	sum := make([]int, l)
	_ = copy(sum, nums)

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if j+1 != i {
				if nums[i]+sum[j] > sum[i] {
					sum[i] = nums[i] + sum[j]
				}
			} else {
				if sum[j] > sum[i] {
					sum[i] = sum[j]
				}
			}
		}
	}
	return sum[l-1]
}

func RobSimplify(nums []int) int {
	prevMax := 0
	curMax := 0
	for _, n := range nums {
		temp := curMax
		if prevMax+n > curMax {
			curMax = prevMax + n
		}
		prevMax = temp
	}
	return curMax
}

func Rob2(nums []int) int {
	var include, exclude int
	for _, n := range nums {
		t, e := include, exclude
		include = exclude + n
		exclude = utils.Max(t, e)
	}
	return utils.Max(include, exclude)
}
