package pk

import "sort"

func CanPK(nums []int, k int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}
	if sum%k == 1 || sum < k {
		return false
	}
	if k == 1 {
		return true
	}
	target := sum / k
	// sort array
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})
	memo := make([]int, k)
	return helper(nums, target, memo, 0)
}

func helper(nums []int, sum int, memo []int, idx int) bool {
	if idx == len(nums) {
		for _, n := range memo {
			if n != sum {
				return false
			}
		}
		return true
	}
	for i := range memo {
		if memo[i]+nums[idx] <= sum {
			memo[i] += nums[idx]
			if helper(nums, sum, memo, idx+1) {
				return true
			}
			memo[i] -= nums[idx]
		}
	}
	return false
}

func helper2(nums []int, visited []bool, start, k, curSum, curIdx, target int) bool {
	if k == 1 {
		return true
	}
	if curSum == target && curIdx > 0 {
		return helper2(nums, visited, 0, k-1, 0, 0, target)
	}
	for i := start; i < len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			if helper2(nums, visited, i+1, k, curSum+nums[i], curIdx+1, target) {
				return true
			}
			visited[i] = false
		}
	}
	return false
}
