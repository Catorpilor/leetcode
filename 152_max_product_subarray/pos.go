package pos

import (
	"github.com/catorpilor/leetcode/utils"
)

func MaxPos(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	maxP := make([]int, n)
	maxP[0] = nums[0]
	ret := maxP[0]
	lastNidx := n
	if nums[0] < 0 {
		lastNidx = 0
	}
	for i := 1; i < n; i++ {
		if nums[i] > 0 || lastNidx == n {
			maxP[i] = utils.Max(nums[i], nums[i]*maxP[i-1])
			if nums[i] < 0 {
				lastNidx = i
			}
		} else {
			tempM, j := 1, i
			for ; j >= lastNidx; j-- {
				tempM *= nums[j]
			}
			maxP[i] = tempM
			if lastNidx > 0 {
				maxP[i] = utils.Max(maxP[i], tempM*maxP[lastNidx-1])
			}
			lastNidx = i
		}
		ret = utils.Max(ret, maxP[i])
	}
	return ret
}
