package min

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func MinSubArrayLen(s int, nums []int) int {
	// brute force
	// O(1) space, O(n^2)
	n := len(nums)
	minL, minLSofar := math.MaxInt32, math.MaxInt32
	for i := 0; i < n; i++ {
		cur := 0
		for j := i; j < n; j++ {
			cur += nums[j]
			if cur >= s {
				minL = j - i + 1
				minLSofar = utils.Min(minL, minLSofar)
				break
			}
		}
	}
	if minLSofar == math.MaxInt32 {
		return 0
	}
	return minLSofar
}

func MinSubArrayLen2(s int, nums []int) int {
	// O(n) time
	var j, sum int
	minLSofar := math.MaxInt32
	for i := range nums {
		sum += nums[i]
		for sum >= s {
			minLSofar = utils.Min(minLSofar, i-j+1)
			sum -= nums[j]
			j++
		}
	}
	if minLSofar == math.MaxInt32 {
		return 0
	}
	return minLSofar
}

func MinSubArrayLen3(s int, nums []int) int {
	// O(nlgn)
	sums := make([]int, len(nums)+1)
	minL := math.MaxInt32
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i < len(sums); i++ {
		if sums[i] >= s {
			p := upper_bound(sums, 0, i, sums[i]-s)
			if p != -1 {
				minL = utils.Min(minL, i-p+1)
			}
		}
	}
	if minL == math.MaxInt32 {
		return 0
	}
	return minL
}

func upper_bound(nums []int, low, high, v int) int {
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] <= v {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if nums[high] > v {
		return high
	}
	return -1
}
