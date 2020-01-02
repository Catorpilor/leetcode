package vtn

import "sort"

func triangleNumbers(nums []int) int {
	return bruteForce(nums)
}

func bruteForce(nums []int) int {
	var res int
	n := len(nums)
	if n < 3 {
		return res
	}
	// sort nums O(nlgn)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	// time complexity O(n^3)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j] > nums[k] {
					res++
				} else {
					break
				}
			}
		}
	}

	return res
}

func binarySearch(nums []int, l, r, x int) int {
	for r >= l && r < len(nums) {
		mid := (l + r) / 2
		if nums[mid] >= x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// use binary search
// time complexity O(n^2lgn)
// space complexity O(1)
func useBs(nums []int) int {
	var res int
	n := len(nums)
	if n < 3 {
		return res
	}
	// sort the slice
	sort.Slice(nums, func(i, j int) bool { return nums[i] <= nums[j] })
	for i := 0; i < n-2; i++ {
		k := i + 2
		for j := i + 1; j < n-1 && nums[i] != 0; j++ {
			// use binarySearch to find the left index that less or equal to nums[i]+nums[j]
			k = binarySearch(nums, k, n-1, nums[i]+nums[j])
			// so from (j+1, k-1) all are valid triangle
			// add that to res
			res += k - j - 1
		}
	}
	return res
}

// func helper(nums []int, res *int, pos int) {

// }
