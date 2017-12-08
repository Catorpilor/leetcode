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
