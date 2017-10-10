package max

import "github.com/catorpilor/leetcode/utils"

func MaxConsecutiveOnes(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var maxSofar, curMax int
	lastFlippedIdx := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			curMax += 1
			maxSofar = utils.Max(maxSofar, curMax)
		} else {
			if lastFlippedIdx == -1 {
				lastFlippedIdx = i
				if i > 0 && nums[i-1] == 1 {
					curMax += 1
					maxSofar = utils.Max(maxSofar, curMax)
				} else {
					maxSofar = utils.Max(maxSofar, curMax)
					curMax = 1
				}
			} else {
				// already flipped
				curMax = i - lastFlippedIdx
				lastFlippedIdx = i
				maxSofar = utils.Max(curMax, maxSofar)
			}
		}
	}
	return utils.Max(maxSofar, 1)
}

func MaxConsecutiveOnes2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	curMax, maxSofar := 1, 0
	lastFlipIdx := -1
	if nums[0] == 0 {
		lastFlipIdx = 0
	}
	for i := 1; i < n; i++ {
		if nums[i] == 1 {
			curMax += 1
		} else {
			if lastFlipIdx == -1 {
				// not flipped
				curMax += 1
			} else {
				// recaculate curMax set the distance between current
				// and lastFlipIdx
				curMax = i - lastFlipIdx
			}
			lastFlipIdx = i
		}
		maxSofar = utils.Max(maxSofar, curMax)
	}
	return maxSofar
}
