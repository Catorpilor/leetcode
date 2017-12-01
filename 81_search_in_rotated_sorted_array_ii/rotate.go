package rotate

func Search(nums []int, target int) bool {
	n := len(nums)
	if n < 1 {
		return false
	}
	minIdx := 0
	if nums[0] >= nums[n-1] {
		for i := 0; i < n-1; i++ {
			if nums[i] > nums[i+1] {
				minIdx = i + 1
			}
		}
	}
	var ret bool
	if target > nums[0] && minIdx > 0 {
		ret = helper(nums, 0, minIdx-1, target)
	} else if target == nums[0] {
		return true
	} else {
		ret = helper(nums, minIdx, n-1, target)
	}
	return ret
}

func helper(nums []int, low, high, target int) bool {
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if nums[low] == target {
		return true
	}
	return false
}

func Search2(nums []int, target int) bool {
	n := len(nums)
	if n < 1 {
		return false
	}
	low, high := 0, n-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[high] {
			if nums[mid] > target && nums[low] <= target {
				high = mid
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] {
			if nums[mid] < target && nums[high] >= target {
				low = mid + 1
			} else {
				high = mid
			}
		} else {
			// nums[mid] == nums[high]
			// ignore the duplicates
			high--
		}
	}
	if nums[low] == target {
		return true
	}
	return false
}

func Search3(nums []int, target int) bool {
	n := len(nums)
	if n < 1 {
		return false
	}
	low, high := 0, n-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[high] {
			if nums[mid] > target && nums[low] <= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] {
			if nums[mid] < target && nums[high] >= target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			// nums[mid] == nums[high]
			// ignore the duplicates
			high--
		}
	}
	return false
}
