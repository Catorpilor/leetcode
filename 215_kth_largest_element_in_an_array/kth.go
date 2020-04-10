package kth

import (
	"math/rand"
	"sort"
)

func kthLargest(nums []int, k int) int {
	return useSort(nums, k)
}

// useSort time complexity O(nlgn), space complexity O(1)
func useSort(nums []int, k int) int {
	n := len(nums)
	if n < 1 || k > n {
		return 0
	}
	sort.Ints(nums)
	return nums[n-k]
}

// useQuickSelect average time compleixty is O(N). worst case is O(N^2)
func useQuickSelect(nums []int, k int) int {
	n := len(nums)
	if n < 1 || k > n {
		return -1
	}
	// same as find (n-k)th smallest in the array
	k = n - k
	l, r := 0, n-1
	for l <= r {
		// random number in [l,r] range
		pivot := rand.Intn(r+1-l) + l // use rand here to prevent the worst case
		// swap pivot and l
		nums[l], nums[pivot] = nums[pivot], nums[l]
		pivot = l
		for j := l + 1; j <= r; j++ {
			if nums[j] < nums[l] {
				// swap
				pivot++
				if j != pivot {
					nums[j], nums[pivot] = nums[pivot], nums[j]
				}
			}
		}
		// swap back now nums[:pivot] all < nums[pivot], nums[pivot:] >= nums[pivot]
		nums[l], nums[pivot] = nums[pivot], nums[l]
		if pivot == k {
			return nums[pivot]
		} else if pivot < k {
			l = pivot + 1
		} else {
			r = pivot - 1
		}
	}
	return -1
}
