package set

import "sort"

func SubSets(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
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
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			n := len(temp)
			temp = append(temp, nums[i])
			helper(res, nums, temp, i+1, target)
			temp = temp[:n]
		}
	}
}

func SubSets2(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	var temp []int
	var ret [][]int
	helper2(nums, 0, &temp, &ret)
	return ret
}

func helper2(nums []int, idx int, temp *[]int, res *[][]int) {
	bak := make([]int, len(*temp))
	copy(bak, *temp)
	*res = append(*res, bak)

	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}
		n := len(*temp)
		*temp = append(*temp, nums[i])
		helper2(nums, i+1, temp, res)
		*temp = (*temp)[:n]
	}
}
