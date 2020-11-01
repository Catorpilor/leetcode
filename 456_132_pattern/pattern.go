package pattern

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func findPattern(nums []int) bool {
	return useBruteForce(nums)
}

// useBruteForce time complexity O(N^3), space complexity O(1)
func useBruteForce(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] < nums[k] && nums[j] > nums[k] {
					return true
				}
			}
		}
	}
	return false
}

// useBetterBF time complexity O(N^2), space complexity O(1)
func useBetterBF(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	// we choose nums[i] as the min value we met so far from [0..j)
	minI := math.MaxInt32
	for j := 0; j < n-1; j++ {
		minI = utils.Min(minI, nums[j])
		for k := j + 1; k < n; k++ {
			if nums[j] > nums[k] && nums[k] > minI {
				return true
			}
		}
	}
	return false
}

// useStack time complexity O(N), space complexity O(N)
func useStack(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	mins := make([]int, n)
	mins[0] = nums[0]
	for i := 1; i < n; i++ {
		mins[i] = utils.Min(mins[i-1], nums[i])
	}
	// st store the possbile value for nums[k]
	st := make([]int, 0, n)
	for j := n - 1; j >= 0; j-- {
		if nums[j] > mins[j] {
			// a possible choice for nums[j]
			nst := len(st)
			for nst > 0 && mins[j] >= st[nst-1] {
				nst--
				st = st[:nst]
			}
			if nst > 0 && st[nst-1] < nums[j] {
				return true
			}
			// a possible nums[k]
			st = append(st, nums[j])
		}
	}
	return false
}
