package max

import (
	"github.com/catorpilor/leetcode/utils"
)

func MaxArea(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	var maxSofar, curMax int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				curMax = nums[j] * (j - i)
			} else {
				curMax = nums[i] * (j - i)
			}
			maxSofar = utils.Max(maxSofar, curMax)
		}
	}
	return maxSofar
}

func MaxArea2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	var maxSofar int
	low, high := 0, n-1
	for low < high {
		maxSofar = utils.Max(maxSofar, (high-low)*utils.Min(nums[low], nums[high]))
		if nums[low] < nums[high] {
			low += 1
		} else {
			high -= 1
		}
	}
	return maxSofar
}
