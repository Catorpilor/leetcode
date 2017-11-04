package ws

import "github.com/catorpilor/leetcode/utils"

func WiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	return 1 + utils.Max(helper(nums, 0, true), helper(nums, 0, false))
}

func helper(nums []int, idx int, flag bool) int {
	var maxLen int
	for i := idx + 1; i < len(nums); i++ {
		if flag && nums[i] < nums[idx] || !flag && nums[i] > nums[idx] {
			maxLen = utils.Max(maxLen, 1+helper(nums, i, !flag))
		}
	}
	return maxLen
}

func WiggleMaxLengthDP1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	// up[i] refers to the length of the longest wiggle subsequence
	// obtained so far considering element as the last element of the wiggle
	// subsequence and ending with a rising wiggle.

	// down[i] refers to the length of the longest wiggle subsequence
	// obtained so far considering element as the last element of the wiggle
	// subsequence and ending with a falling wiggle.
	up, down := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				up[i] = utils.Max(up[i], down[j]+1)
			} else if nums[j] > nums[i] {
				down[i] = utils.Max(down[i], up[j]+1)
			}
		}
	}
	return 1 + utils.Max(up[n-1], down[n-1])
}

func WiggleMaxLengthDP2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	// up[i] refers to the length of the longest wiggle subsequence
	// obtained so far considering element as the last element of the wiggle
	// subsequence and ending with a rising wiggle.

	// down[i] refers to the length of the longest wiggle subsequence
	// obtained so far considering element as the last element of the wiggle
	// subsequence and ending with a falling wiggle.
	up, down := make([]int, n), make([]int, n)
	up[0], down[0] = 1, 1

	// for every elemnt in nums it could be one of these three category.
	// a. nums[i] > nums[i-1] means the rising wiggle, so the element before
	// must be a falling position
	// so we have up[i] = down[i-1]+1, down[i] = down[i-1]
	// b. nums[i] < nums[i-1] similarly we can have
	// down[i] = up[i-1]+1, up[i] = up[i-1]
	// c nums[i] == nums[i-1]
	// up[i] = up[i-1] down[i] = down[i-1]
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up[i], down[i] = down[i-1]+1, down[i-1]
		} else if nums[i] < nums[i-1] {
			down[i], up[i] = up[i-1]+1, up[i-1]
		} else {
			down[i], up[i] = down[i-1], up[i-1]
		}
	}
	return utils.Max(up[n-1], down[n-1])
}

func WiggleMaxLengthDP3(nums []int) int {
	// optimized space based on dp2
	// case up[i], down[i] only related to up[i-1],down[i-1]
	n := len(nums)
	if n < 2 {
		return n
	}
	up, down := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	return utils.Max(up, down)
}

func WiggleMaxLengthGreedy(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	prevDiff := nums[1] - nums[0]
	ret := 1
	if prevDiff != 0 {
		ret = 2
	}
	for i := 2; i < n; i++ {
		df := nums[i] - nums[i-1]
		if prevDiff <= 0 && df > 0 || prevDiff >= 0 && df < 0 {
			ret += 1
			prevDiff = df
		}
	}
	return ret
}
