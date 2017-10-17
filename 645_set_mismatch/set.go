package set

func FindNums(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nil
	}
	// allocate a set slice
	set := make([]int, n+1)
	ret := make([]int, 0, 2)
	for _, c := range nums {
		set[c] += 1
		if set[c] > 1 {
			ret = append(ret, c)
		}
	}
	for i, c := range set {
		if c == 0 && i != 0 {
			ret = append(ret, i)
			break
		}
	}
	return ret
}

func FindNums2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	ret := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ret = append(ret, nums[i], i+1)
			break
		}
	}
	return ret

}
