package sq

import (
	"github.com/catorpilor/leetcode/utils"
)

func LengthOfLCIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	maxSofar, curMax := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			curMax += 1
			maxSofar = utils.Max(maxSofar, curMax)
		} else {
			curMax = 1
		}
	}
	return maxSofar
}
