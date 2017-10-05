package num

func FindDisappearedNumbers(nums []int) []int {
	if len(nums) <= 1 {
		return nil
	}
	// with extra space
	h := make([]int, len(nums)+1)
	for _, c := range nums {
		h[c] += 1
	}
	var ret []int
	for i := 1; i < len(h); i++ {
		if h[i] == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}

func findDisappearedNumbers(nums []int) []int {
	if len(nums) <= 1 {
		return nil
	}
	var val int
	// a map between array idx[0,n-1] to nums [1,n]
	// for example [3,2,2]
	// i=0, nums[i] = 3,
	// val = 2 nums[val] = 2 > 0
	// so we flip nums[val] = -2
	// which tells us that 2 has been seen in the array
	for i := 0; i < len(nums); i++ {
		val = abs(nums[i]) - 1
		if nums[val] > 0 {
			nums[val] = -nums[val]
		}
	}
	var ret []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	if a == 0 {
		return 0
	}
	return a
}
