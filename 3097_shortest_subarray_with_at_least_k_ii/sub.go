package sub

import (
	"github.com/Catorpilor/leetcode/utils"
)

func shortestSubArrWithAtLeastK(nums []int, k int) int {
	n := len(nums)
	acc, rcc := make([]int, n), make([]int, n)
	acc[0] = nums[0]
	rcc[n-1] = nums[n-1]
	if nums[0] >= k || nums[n-1] >= k {
		return 1
	}
	lcur := -1
	for i := 1; i < n; i++ {
		if nums[i] >= k {
			return 1
		}
		acc[i] = acc[i-1] | nums[i]
		if acc[i] >= k && lcur == -1 {
			lcur = i
		}
	}
	if lcur == -1 {
		return -1
	}

	rcur := -1
	for i := n - 2; i >= 0; i-- {
		rcc[i] = rcc[i+1] | nums[i]
		if rcc[i] >= k && rcur == -1 {
			rcur = i
		}
	}
	if rcur == -1 {
		return -1
	}
	// fmt.Printf("lcur:%d, rcur:%d\n", lcur, rcur)

	return utils.Min(lcur+1, n-rcur)
}

func useSlidingWindow(nums []int, k int) int {
	n := len(nums)
	acc := make([]int, n)
	acc[0] = nums[0]
	l, r := 0, n+1
	for i := 1; i < n; i++ {
		acc[i] = acc[i-1] + nums[i]
		if acc[i] >= k {
			r = utils.Min(r, i)
		}
	}
}
