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
