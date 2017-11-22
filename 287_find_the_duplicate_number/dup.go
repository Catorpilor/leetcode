package dup

import (
	"sort"
)

func FindDuplicate(nums []int) int {
	// brute forced
	n := len(nums)
	if n <= 1 {
		return -1
	}
	// O(n^2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1

}

func FindDuplicate2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return -1
	}
	// sort nums
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

func FindDuplicate3(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return -1
	}
	/* mathematic proof
	* assume at the
	* meeting point slow moves s steps, fast moves 2s steps, and the lenght of circle is c
	* then we have 2s = s + n * c -> s = n * c -- (1)
	* then assume the length from start point to entry point is x
	* and the length from entry point to meeting point is a
	* there must be s = x + a , so
	* x + a = s = n * c
	* x = (n-1)*c + (c-a) // c-a means the length from meeting point to entry point
	* LHS(x) means the fast tag moves x steps
	* RHS means slow moves (n-1) cycles plus the length from the meeting point to entry point
	* so we get the entry point
	 */
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	// rest fast to 0
	fast = 0
	for fast != slow {
		fast, slow = nums[fast], nums[slow]
	}
	return slow
}

func FindDuplicate4(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return -1
	}
	low, high, mid := 1, n-1, 0
	for low < high {
		mid = low + (high-low)/2
		count := 0
		for _, v := range nums {
			if v <= mid {
				count += 1
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low

}
