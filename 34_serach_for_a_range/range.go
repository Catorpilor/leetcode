package ranger

func SearchRange(nums []int, target int) []int {
	n := len(nums)
	ret := []int{-1, -1}
	if n < 1 || n == 1 && nums[0] != target || nums[0] > target || nums[n-1] < target {
		return ret
	}
	ret[0] = lower_bound(nums, 0, n-1, target)
	if ret[0] == -1 {
		return ret
	}
	ret[1] = upper_bound(nums, ret[0], n-1, target)
	return ret
}

func upper_bound(nums []int, low, high, target int) int {
	for low < high {
		mid := (high + low + 1) / 2 //trick part to use high-based mid
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid
		}
	}
	if nums[high] == target {
		return high
	}
	return -1
}
func lower_bound(nums []int, low, high, target int) int {
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if nums[low] == target {
		return low
	}
	return -1
}

func SearchRange2(nums []int, target int) []int {
	n := len(nums)
	ret := []int{-1, -1}
	if n < 1 || nums[0] > target || nums[n-1] < target {
		return ret
	}
	start := lower_bound2(nums, 0, n, target)
	if start == n || nums[start] != target {
		return ret
	}
	ret[0] = start
	ret[1] = lower_bound2(nums, 0, n, target+1) - 1
	return ret
}

func lower_bound2(nums []int, low, high, target int) int {
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
