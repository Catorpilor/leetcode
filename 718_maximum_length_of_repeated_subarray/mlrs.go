package mlrs

import (
	"github.com/catorpilor/leetcode/utils"
)

func MaxLength(nums1, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return 0
	}
	m, n := n1, n2
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	var ret int
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				ret = utils.Max(ret, dp[i][j])
			}
		}
	}
	return ret
}
