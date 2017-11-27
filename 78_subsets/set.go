package set

func SubSets(nums []int) [][]int {
	ret := make([][]int, 0)
	ret = append(ret, []int{}, nums)
	for i := 1; i < len(nums); i++ {
		temp := make([]int, 0)
		helper(&ret, nums, temp, 0, i)
	}
	return ret
}

func helper(res *[][]int, nums, temp []int, start, target int) {
	if len(temp) == target {
		bak := make([]int, target)
		copy(bak, temp)
		*res = append(*res, bak)
	} else {
		for i := start; i < len(nums); i++ {
			n := len(temp)
			temp = append(temp, nums[i])
			helper(res, nums, temp, i+1, target)
			temp = temp[:n]
		}
	}
}

func SubSets2(nums []int) [][]int {
	// bit manupulation
	n := len(nums)
	total := 1 << uint(n)
	ret := make([][]int, 0)
	for i := 0; i < total; i++ {
		var cur []int
		for j := 0; j < n; j++ {
			if (i>>uint(j))&1 != 0 {
				cur = append(cur, nums[j])
			}
		}
		ret = append(ret, cur)
	}
	return ret
}
