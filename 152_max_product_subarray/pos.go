package pos

import (
	"github.com/catorpilor/leetcode/utils"
)

func MaxPos(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	// maxP stores the product for ith element in nums
	// ex : nums = [7,-2,-4]
	// related maxP = [7,-2,56]
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
			// travese until the first negtive number occurs
			// set lastIdx to i
			if nums[i] < 0 {
				lastNidx = i
			}
		} else {
			tempM, j := 1, i
			// go back to the lastNidx
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

func MaxPos2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	// maxsoFar stores the max product we've found
	// so far
	maxSofar := nums[0]
	for i, maxv, minv := 1, maxSofar, maxSofar; i < len(nums); i++ {
		// multipled by a negative number makes a big number smaller,
		// small number bigger, so we swap them
		if nums[i] < 0 {
			maxv, minv = minv, maxv
		}
		maxv = utils.Max(nums[i], maxv*nums[i])
		minv = utils.Min(nums[i], minv*nums[i])

		maxSofar = utils.Max(maxSofar, maxv)
	}
	return maxSofar

}
