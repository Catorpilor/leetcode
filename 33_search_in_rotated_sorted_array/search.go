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
