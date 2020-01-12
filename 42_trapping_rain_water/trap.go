package trap

import (
	"github.com/catorpilor/leetcode/utils"
)

func Trap(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return 0
	}
	// same as largest ractangle in histogram
	// use stack to store indexes.
	s := utils.NewStack()
	var i, ret int
	for i < n {
		if s.IsEmpty() || nums[s.Top().(int)] >= nums[i] {
			s.Push(i)
			i++
		} else {
			v := s.Pop().(int) // bottom
			// nums[i] represents the right boundary

			if !s.IsEmpty() { // left boundary found
				top := s.Top().(int)
				// nums[top] represents the left bounday
				// we have to find the minimal one
				ret += (utils.Min(nums[top], nums[i]) - nums[v]) * (i - top - 1)
			}
		}
	}
	return ret
}

func Trap2(nums []int) int {
	// O(n) time and O(1) space
	n := len(nums)
	if n <= 2 {
		return 0
	}
	left, right, maxLeft, maxRight := 0, n-1, 0, 0
	var res int
	// Search from left to right and maintain a max height of left and right separately,
	// which is like a one-side wall of partial container.
	// Fix the higher one and flow water from the lower part.
	// For example, if current height of left is lower, we fill water in the left bin.
	// Until left meets right, we filled the whole container.
	for left <= right {
		if nums[left] <= nums[right] {
			if nums[left] >= maxLeft {
				maxLeft = nums[left]
			} else {
				res += maxLeft - nums[left]
			}
			left++
		} else {
			if nums[right] >= maxRight {
				maxRight = nums[right]
			} else {
				res += maxRight - nums[right]
			}
			right--
		}
	}
	return res
}

func bruteForce(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return 0
	}
	var res, maxL, maxR int
	for i := range nums {
		maxL, maxR = 0, 0
		for j := 0; j < i; j++ {
			if nums[j] > maxL {
				maxL = nums[j]
			}
		}
		for j := i + 1; j < n; j++ {
			if nums[j] > maxR {
				maxR = nums[j]
			}
		}
		units := utils.Min(maxL, maxR) - nums[i]
		if units > 0 {
			res += units
		}

	}
	return res
}

// dp same idea as bruteForce, but use two slice to store the max height
// from both directions, time complexity is 0(N), space complexity is O(N)
func dp(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return 0
	}
	maxL := make([]int, n)
	maxL[0] = nums[0]
	for i := 1; i < n; i++ {
		maxL[i] = utils.Max(nums[i], maxL[i-1])
	}
	maxR := make([]int, n)
	maxR[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		maxR[i] = utils.Max(nums[i], maxR[i+1])
	}
	var res int
	for i := 1; i < n-1; i++ {
		res += utils.Min(maxL[i], maxR[i]) - nums[i]
	}
	return res
}
