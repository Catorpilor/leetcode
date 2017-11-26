package peak

func PeakIdx(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	for i := 0; i < n-1; i++ {
		if i > 0 && i < n-1 {
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				return i
			}
		} else {
			// i ==0
			if nums[1] < nums[0] {
				return 0
			}
			if nums[n-1] > nums[n-2] {
				return n - 1
			}
		}
	}
	return -1
}

func PeakIdx2(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			return i - 1
		}
	}
	return n - 1
}

func PeakIdx3(nums []int) int {
	// binary search
	// based on the condition nums[i] != nums[i+1]
	// so if nums[mid]<nums[mid+1] then peak might in nums[mid+1:]
	// otherwise peak might in nums[0:mid]
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < nums[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
