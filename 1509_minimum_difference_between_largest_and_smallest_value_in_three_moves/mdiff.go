package mdiff

import (
	"math"
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func minDiffs(nums []int) int {
	return useSort(nums)
}

// useSort time complexity O(nlogn), space compleixty O(1)
func useSort(nums []int) int {
	n := len(nums)
	if n < 5 {
		return 0
	}
	sort.Ints(nums)
	// 4 cases
	// 1. remove 3 big nums --> nums[n-4] - nums[0]
	// 2. remove 1 smallest, 2 bigest nums nums[n-3] - nums[1]
	// 3. remove 2 smallest, 1 big nums nums[n-2] - nums[2]
	// 4. remove 3 smallest. nums[n-1]- nums[3]
	res := math.MaxInt32
	for i := 0; i < 4; i++ {
		res = utils.Min(res, nums[n-4+i]-nums[i])
	}
	return res
}
