package targetsum

import "sort"

func maxSubarray(nums []int, target int) int {
	return useHashmap(nums, target)
}

func useHashmap(nums []int, target int) int {
	n := len(nums)
	set := make(map[int]int, n)
	set[0] = -1
	sum := 0
	var ans int
	for i := 0; i < n; i++ {
		sum += nums[i]
		if _, exists := set[sum-target]; exists {
			ans++
			clear(set)
		}
		set[sum] = i
	}
	return ans
}

func clear(set map[int]int) {
	for k := range set {
		delete(set, k)
	}
}

// useHashmapAndGreedy time complexity O(N*Log(N)), space complexity O(N)
func useHashmapAndGreedy(nums []int, target int) int {
	n := len(nums)
	set := make(map[int]int, n)
	// key -> sum, val -> start index.
	set[0] = -1
	// pairs store the subarraies with the sum equals the target
	pairs := make([][]int, 0, n)
	var sum int
	for i := 0; i < n; i++ {
		sum += nums[i]
		if v, exists := set[sum-target]; exists {
			pairs = append(pairs, []int{v + 1, i})
		}
		set[sum] = i
	}
	np := len(pairs)
	if np <= 1 {
		return np
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	ans := 1
	end := pairs[0][1]
	for i := 1; i < np; i++ {
		if pairs[i][0] > end {
			end = pairs[i][1]
			ans++
		}
	}
	return ans
}
