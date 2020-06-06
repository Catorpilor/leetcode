package subset

import "sort"

func largestSubset(nums []int) []int {
	return useDP(nums)
}

// useDP time complexity O(N^2), space complexity O(N)
func useDP(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	// sort here to avoid cases like [6,2,8]
	// nums[i] % nums[j] == 0 || nums[j] % nums[i] == 0
	// will return 3 but actually the largest subset it 2.
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	count := make([]int, n)
	pres := make([]int, n)
	max, index := 0, -1
	for i := 0; i < n; i++ {
		count[i] = 1
		pres[i] = -1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if count[j]+1 > count[i] {
					count[i] = count[j] + 1
					pres[i] = j
				}
			}
		}
		if count[i] > max {
			max = count[i]
			index = i
		}
	}
	ans := make([]int, 0, n)
	for index != -1 {
		ans = append(ans, nums[index])
		index = pres[index]
	}
	return ans
}
