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

func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	set := make(map[int]bool, n)
	backtrack(&ans, []int{}, nums, set, n)
	return ans
}

// backtrack time complexity O(N*N!), space complexity O(N)
func backtrack(ans *[][]int, store, nums []int, set map[int]bool, n int) {
	ns := len(store)
	if ns == n {
		*ans = append((*ans), store)
		return
	}
	for i := 0; i < n; i++ {
		if !set[nums[i]] {
			store = append(store, nums[i])
			set[nums[i]] = true
			backtrack(ans, store, nums, set, n)
			set[nums[i]] = false
			store = store[:ns]
		}
	}
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

func permuteCommon(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	set := make(map[int]int, n)
	for i := range nums {
		set[nums[i]]++
	}
	useBF(&ans, set, []int{}, nums, n)
	return ans
}

// useBF the same approach to solve 47
func useBF(ans *[][]int, set map[int]int, cur, nums []int, n int) {
	if n == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*ans = append((*ans), tmp)
		return
	}
	for k := range set {
		if set[k] > 0 {
			nc := len(cur)
			set[k]--
			cur = append(cur, k)
			useBF(ans, set, cur, nums, n-1)
			set[k]++
			cur = cur[:nc]
		}
	}
	return
}
