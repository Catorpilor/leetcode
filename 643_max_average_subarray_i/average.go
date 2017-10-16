package average

import (
	"fmt"

	"github.com/catorpilor/leetcode/utils"
)

func MaxAverage(nums []int, k int) float64 {
	if len(nums) < k || k == 0 {
		return float64(0)
	}
	var prev, cur, maxSofar int
	for i := 0; i < k; i++ {
		prev += nums[i]
	}
	maxSofar = prev
	for i := 1; i < len(nums)-k+1; i++ {
		cur = prev - nums[i-1] + nums[i+k-1]
		maxSofar = utils.Max(cur, maxSofar)
		prev = cur
	}
	fmt.Println(maxSofar)
	return float64(maxSofar) / float64(k)
}

func MaxAverage2(nums []int, k int) float64 {
	// cumulative sum array
	// sums[i] stores the sum of the elements of the
	// given nums array from 0 to ith index
	if len(nums) < k || k == 0 {
		return float64(0)
	}
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	ret := sums[k-1]
	for i := k; i < len(sums); i++ {
		ret = utils.Max(ret, sums[i]-sums[i-k])
	}
	return float64(ret) / float64(k)
}

func MaxAverage3(nums []int, k int) float64 {
	// sliding window
	// same as MaxAverage
	if len(nums) < k || k == 0 {
		return float64(0)
	}
	var cur, maxSofar int
	for i := 0; i < k; i++ {
		cur += nums[i]
	}
	maxSofar = cur
	for i := k; i < len(nums); i++ {
		cur += nums[i] - nums[i-k]
		maxSofar = utils.Max(cur, maxSofar)
	}
	return float64(maxSofar) / float64(k)
}
