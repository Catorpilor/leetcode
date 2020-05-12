package single

func singleDup(nums []int) int {
	return useBinarySearch(nums)
}

// useBinarySearch time complexity O(lgN), space complexity O(1)
func useBinarySearch(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
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

// useBSSimple time complexity O(lgN), space complexity O(1)
func useBSSimple(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		// mid^1 is brilliant and elegent.
		// for example with nums = [1,1,2,2,3]
		//                          0,1,2,3,4
		// mid = 2, mid^1=3
		// nums[mid] == nums[mid^1] then l = 3
		// mid = 3, mid^i = 2
		// nums[mid] == nums[mid^1] then i=4 break the loop.
		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
