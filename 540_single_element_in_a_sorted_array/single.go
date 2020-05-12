package single

func singleDup(nums []int) int {
	return useBinarySearch(nums)
}

func useBinarySearch(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		if l == r {
			break
		}
		mid := l + (r-l)/2
		if mid > 0 && mid < n-1 && nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			// find the one
			return nums[mid]
		} else if mid > 0 && nums[mid] != nums[mid-1] {
			if mid%2 == 0 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else if mid < n-1 && nums[mid] != nums[mid+1] {
			if (mid+1)%2 == 0 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return nums[l]
}
