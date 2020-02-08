package pot

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func MaxPot(nums []int) int {
	// sort slice
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	n := len(nums)
	ret := nums[n-1] * nums[n-2] * nums[n-3]
	ret = utils.Max(ret, nums[0]*nums[1]*nums[n-1])
	return ret
}

type queue struct {
	max1, max2, max3, min1, min2 int
}

func (q *queue) Append(n int) {
	if q.max1 < n {
		q.max3 = q.max2
		q.max2 = q.max1
		q.max1 = n
	} else if q.max2 < n {
		q.max3 = q.max2
		q.max2 = n
	} else if q.max3 < n {
		q.max3 = n
	}

	if q.min1 > n {
		q.min2 = q.min1
		q.min1 = n
	} else if q.min2 > n {
		q.min2 = n
	}
}

func MaxPot2(nums []int) int {
	q := &queue{
		max1: -1000,
		max2: -1000,
		max3: -1000,
		min1: 1000,
		min2: 1000,
	}
	for _, c := range nums {
		q.Append(c)
	}
	ret := q.max1 * q.max2 * q.max3
	ret = utils.Max(ret, q.min1*q.min2*q.max1)
	return ret
}
