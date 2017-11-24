package sum

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func ThreeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n <= 3 {
		return sum(nums)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if utils.Abs(target-ans) > utils.Abs(target-sum) {
				ans = sum
				if ans == target {
					return target
				}
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}

func sum(nums []int) int {
	var ret int
	for i := range nums {
		ret += nums[i]
	}
	return ret
}
