package max

func MaxConsecutiveOnes(nums []int) int {
	var curMax, maxSofar int
	for _, c := range nums {
		if c == 1 {
			curMax += 1
		} else {
			maxSofar = max(maxSofar, curMax)
			curMax = 0
		}
	}
	return max(maxSofar, curMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
