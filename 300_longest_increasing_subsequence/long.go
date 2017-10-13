package long

import (
	"github.com/catorpilor/leetcode/utils"
)

func LengthOfLIS(nums []int) int {
	// DP
	// O(n^2) time
	// O(1) space
	n := len(nums)
	if n < 2 {
		return n
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	var ret int
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = utils.Max(dp[i], 1+dp[j])
			}
			ret = utils.Max(ret, dp[i])
		}
	}
	return ret
}

func LengthOfLIS2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	lis := make([]int, 0, n)
	for _, n := range nums {
		pos := binarySearch(lis, n)
		if pos == len(lis) {
			lis = append(lis, n)
		}
		if lis[pos] > n {
			lis[pos] = n
		}
	}
	return len(lis)
}

func binarySearch(nums []int, v int) int {
	// nums is sorted in ascending order
	b, e := 0, len(nums)
	for b < e {
		mid := b + (e-b)>>1
		if nums[mid] > v {
			e = mid
		} else if nums[mid] == v {
			return mid
		} else {
			b = mid + 1
		}
	}
	return b
}
