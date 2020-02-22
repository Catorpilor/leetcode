package suca

import "fmt"

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
