package jump

import (
	"math"
)

func Jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[n-1], dp[n-2] = 0, 1
	for i := n - 3; i >= 0; i-- {
		if nums[i] >= n-i-1 {
			dp[i] = 1
		} else {
			for j := 1; j <= nums[i]; j++ {
				if 1+dp[i+j] < dp[i] {
					dp[i] = 1 + dp[i+j]
				}
			}
		}
	}
	return dp[0]
}

func Jump2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	// assume we can always jump to the last index
	var start, end, step, maxend int
	for end < n-1 {
		step++
		maxend = end + 1
		for i := start; i <= end; i++ {
			if i+nums[i] >= n-1 {
				return step
			}
			if i+nums[i] > maxend {
				maxend = i + nums[i]
			}
		}
		start = end + 1
		end = maxend
	}
	return step
}

func Jump3(nums []int) int {
	i, step, end, maxend, n := 0, 0, 0, 0, len(nums)
	for end < n-1 {
		step++
		for ; i <= end; i++ {
			if i+nums[i] > maxend {
				maxend = i + nums[i]
			}
			if maxend >= n-1 {
				return step
			}
		}
		if end == maxend {
			// can not reach the end
			break
		}
		end = maxend
	}
	if n == 1 {
		return 0
	}
	return -1
}
