package search

func IsExist(nums []int, target int) int {
	n := len(nums)
	if n < 1 {
		return -1
	}
	low, high := 0, n-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	rotIdx := low
	low, high = 0, n-1
	for low <= high {
		mid := low + (high-low)/2
		realmid := (mid + rotIdx) % n
		if nums[realmid] == target {
			return realmid
		} else if nums[realmid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	return useBinarySearchWithOnePass(nums, target)
}

// useBinarySearchWithOnePass time complexity O(lgN), space complexity O(1)
func useBinarySearchWithOnePass(nums []int, target int) int {
	n := len(nums)
	if n < 1 {
		return -1
	}
	if target > nums[0] && target < nums[n-1] {
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// useBinarySearchWithTwoPasses time complexity O(lgN), space complexity O(1)
func useBinarySearchWithTwoPasses(nums []int, target int) int {
	n := len(nums)
	if n < 1 || (target > nums[0] && target < nums[n-1]) {
		return -1
	}
	// the following piece of code is to find the real start
	// which means the smallest element in the array.
	// for example
	// with rotated nums:                         [4,5,6,7,0,1,2], we can split the array into two parts.
	// with condition nums[i] > nums[n-1],we got   T,T,T,T,F,F,F
	// so the mission here is to find the first `F`
	initStart := 0
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[n-1] {
			left = mid + 1
		} else {
			initStart = mid
			right = mid - 1
		}
	}
	// the second pass
	left, right = 0, n-1
	if target > nums[n-1] {
		right = initStart - 1
	}
	if target < nums[0] {
		left = initStart
	}
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
