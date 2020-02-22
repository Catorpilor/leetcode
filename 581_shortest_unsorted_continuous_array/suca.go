package suca

import (
	"fmt"
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func unsortedSubarray(nums []int) int {
	return myWay(nums)
}

// myWay time complexity O(N!)
func myWay(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	l, r := helper(nums)
	fmt.Printf("nums: %v, l: %d, r:%d \n", nums, l, r)
	if l != 0 && r != n-1 {
		return n
	} else if l == 0 && r == n-1 {
		return myWay(nums[l+1 : r])
	}
	if l == 0 {
		return myWay(nums[l+1:])
	}
	return myWay(nums[:r])
}

// helper returns nums's min & max position, time complexity O(N)
func helper(nums []int) (int, int) {
	n := len(nums)
	if n <= 1 {
		return 0, 0
	}
	l, r := 0, n-1
	minIdx, maxIdx := l, r
	for l <= r {
		if nums[l] < nums[minIdx] {
			minIdx = l
		}
		if nums[l] > nums[maxIdx] {
			maxIdx = l
		}
		if nums[r] < nums[minIdx] {
			minIdx = r
		}
		if nums[r] > nums[maxIdx] {
			maxIdx = r
		}
		l++
		r--
	}
	return minIdx, maxIdx
}

// selectionSort time complexity O(N^2), space complexity O(1)
func selectionSort(nums []int) int {
	n := len(nums)
	// just like selection sort, l, r are boundaries of the range needed to be sorted
	l, r := n, 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[i] {
				// nums[i] and nums[j] are in the wrong position
				// update l, left most position
				l = utils.Min(l, i)
				// update r, right most position
				r = utils.Max(r, j)
			}
		}
	}
	if r < l {
		return 0
	}
	return r - l + 1
}

// useSort time complexity O(NlgN), space complexity O(N)
func useSort(nums []int) int {
	n := len(nums)
	local := make([]int, n)
	copy(local, nums)
	sort.Slice(local, func(i, j int) bool {
		return local[i] <= local[j]
	})
	l, r := n, 0
	for i := 0; i < n; i++ {
		if local[i] != nums[i] {
			l = utils.Min(l, i)
			r = utils.Max(r, i)
		}
	}
	if l > r {
		return 0
	}
	return r - l + 1
}
