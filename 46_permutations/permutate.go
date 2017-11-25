package permutate

import (
	"sort"
)

func Permute(nums []int) [][]int {
	n := len(nums)

	if n <= 1 {
		return [][]int{nums}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ret := make([][]int, 0, n*(n-1))
	org := make([]int, n)
	copy(org, nums)

	ret = append(ret, org)
	for {
		// ret = append(ret, nums)
		if res, ok := helper(nums); !ok {
			break
		} else {
			ret = append(ret, res)
		}
	}
	return ret
}

func helper(nums []int) ([]int, bool) {
	n := len(nums)
	i := n - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}
	if i <= 0 {
		return nil, false
	}
	j := n - 1
	for nums[j] <= nums[i-1] {
		j--
	}
	nums[i-1], nums[j] = nums[j], nums[i-1]
	j = n - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i, j = i+1, j-1
	}

	temp := make([]int, n)
	copy(temp, nums)
	return temp, true
}

func Permute2(nums []int) [][]int {
	res := make([][]int, 0)
	helper2(nums, 0, len(nums), &res)
	return res
}

func helper2(nums []int, a, b int, res *[][]int) {
	if a == b {
		num_copy := make([]int, len(nums))
		copy(num_copy, nums)
		*res = append(*res, num_copy)
		return
	}
	for i := a; i < b; i++ {
		nums[a], nums[i] = nums[i], nums[a]
		helper2(nums, a+1, b, res)
		nums[a], nums[i] = nums[i], nums[a]
	}
	return
}
