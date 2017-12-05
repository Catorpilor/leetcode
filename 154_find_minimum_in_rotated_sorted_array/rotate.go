package rotate

func FindMinimum(nums []int) int {
	// O(n) time, and O(1) space
	n := len(nums)
	if n < 1 {
		return 0
	}
	var minIdx int
	for i := 1; i < n; i++ {
		if nums[i] < nums[minIdx] {
			minIdx = i
		}
	}
	return nums[minIdx]
}

func FindMinimum2(nums []int) int {
	// O(lgn) time
	n := len(nums)
	if n < 1 {
		return 0
	}
	low, hi := 0, n-1
	for low < hi {
		mid := low + (hi-low)/2
		if nums[mid] > nums[hi] {
			low = mid + 1
		} else if nums[mid] < nums[hi] {
			hi = mid
		} else {
			hi--
		}
	}
	return nums[low]
}
