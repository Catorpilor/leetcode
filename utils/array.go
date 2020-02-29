package utils

// LowerBound returns the lowest pos for target in the sorted array
func LowerBound(nums []int, target int) int {
	n := len(nums)
	low, hi := 0, n-1
	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] >= target {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
