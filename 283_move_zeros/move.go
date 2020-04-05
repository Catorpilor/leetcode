package move

func MoveZeroes(nums []int) []int {
	curZeroIdx := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if curZeroIdx == -1 {
				curZeroIdx = i
			}
		} else {
			if curZeroIdx != -1 {
				nums[curZeroIdx] = nums[i]
				nums[i] = 0
				if curZeroIdx+1 < i {
					curZeroIdx += 1
				} else {
					curZeroIdx = i
				}
			}
		}
	}
	return nums
}

// moveZeros time complexity O(N), space complexity O(1)
func moveZeros(nums []int) []int {
	return useTwoPointers(nums)
}

func MoveZeroes2(nums []int) []int {
	leftZeroIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i > leftZeroIdx {
				nums[leftZeroIdx] = nums[i]
				nums[i] = 0
			}
			leftZeroIdx += 1
		}
	}
	return nums
}

func useTwoPointers(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	left, right := -1, -1 // left, right boundary of zeros, [left, right) are all zeros.
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			continue
		}
		left = i
		if right == -1 {
			right = left + 1
		}
		for right < n && nums[right] == 0 {
			right++
		}
		if right == n {
			break
		}
		// swap left, right
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}

// useOnePointer time complexity O(N), space complexity O(1)
func useOnePointer(nums []int) []int {
	var nxt int // nxt point to the next position for a non-zero value
	for _, num := range nums {
		if num != 0 {
			nums[nxt] = num
			nxt++
		}
	}
	for i := nxt; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}
