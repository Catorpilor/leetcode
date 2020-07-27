package pivot

func findMin(nums []int) int {
	return useBinarySearch(nums)
}

func useStraight(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return nums[0]
}

func useBinarySearch(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	low, high := 0, n-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < nums[high] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return nums[low]
}

// useBS time complexity O(logN), space complexity O(1)
func useBS(nums []int) int {
	// nums 3,4,5,6,0,1,2
	// fmt: T T T T F F F
	// we need to find the first false,
	// so we could use binary search
	n := len(nums)
	if n < 1 {
		return 0
	}
	low, hi := 0, n-1
	var ans int
	for low <= hi {
		mid := low + (hi-low)/2
		if nums[mid] >= nums[0] {
			low = mid + 1
		} else {
			ans = mid
			hi = mid - 1
		}
	}
	return nums[ans]
}
