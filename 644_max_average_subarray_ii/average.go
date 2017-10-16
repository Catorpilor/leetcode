package average

import (
	"fmt"
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func MaxAverage(nums []int, k int) float64 {
	if len(nums) < k || k == 0 {
		return float64(0)
	}
	// cumulative sum array
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	sumk := make([]int, len(nums)-k+1)
	for i := 0; i < len(sumk); i++ {
		sumk[i] = sums[i+k-1]
		for j := k + i; j < len(sums); j++ {
			sumk[i] = utils.Max(sumk[i], sums[j]-sums[j-k-i])
		}
	}
	fmt.Println(sumk)
	ret := float64(sumk[0]) / float64(k)
	for i := 1; i < len(sumk); i++ {
		ret = utils.Maxf(ret, float64(sumk[i])/float64(k+i))
	}
	return ret
}

func MaxAverage2(nums []int, k int) float64 {
	if len(nums) < k || k == 0 {
		return float64(0)
	}
	// binary search
	maxV, minV := float64(nums[0]), float64(nums[0])
	for _, v := range nums {
		if float64(v) > maxV {
			maxV = float64(v)
		}
		if float64(v) < minV {
			minV = float64(v)
		}
	}
	prev_mid, accuracy := float64(maxV), math.MaxFloat64
	for accuracy > 0.00001 {
		mid := float64(maxV+minV) * 0.5
		if check(nums, mid, k) {
			minV = mid
		} else {
			maxV = mid
		}
		accuracy = math.Abs(prev_mid - mid)
		prev_mid = mid
	}
	return minV
}

func check(nums []int, mid float64, k int) bool {
	var sum, prev, min_sum float64
	for i := 0; i < k; i++ {
		sum += float64(nums[i]) - mid
	}
	if sum >= 0 {
		return true
	}
	for i := k; i < len(nums); i++ {
		sum += float64(nums[i]) - mid
		prev += float64(nums[i-k]) - mid
		min_sum = math.Min(min_sum, prev)
		if sum >= min_sum {
			return true
		}
	}
	return false
}
