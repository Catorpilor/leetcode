package pos

func FirstMissingPositive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}
	// not constant space
	hm := make(map[int]bool)
	minV := 1
	for _, v := range nums {
		hm[v] = true
		if v > 0 && v < minV {
			minV = v
		}
	}
	if minV > 1 {
		return 1
	}

	for hm[minV] {
		minV++
	}
	return minV
}

func FirstMissingPositive2(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
