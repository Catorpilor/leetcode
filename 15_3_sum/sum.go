package sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	return useHashMap(nums)
}

// useHashMap time complexity O(N^2), space complexity O(N)
func useHashMap(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if nums[0] > 0 || nums[n-1] < 0 {
		return [][]int{}
	}
	ans := make([][]int, 0, n)
	set := make(map[int][]int, n)
	for i := 0; i <= n-3; i++ {
		if nums[i] > 0 {
			break
		}
		clear(set)
		// nums[i] > nums[i-1] to remote the duplicates
		if i == 0 || nums[i] > nums[i-1] {
			for j := i + 1; j < n; j++ {
				// fmt.Printf("i:%d, j:%d, nums[j]:%d and set: %v\n", i, j, nums[j], set)
				if tmp, exists := set[nums[j]]; exists {
					ans = append(ans, []int{tmp[0], tmp[1], nums[j]})
					for j < n-1 && nums[j] == nums[j+1] {
						j++
					}
					continue
				}
				cur := nums[i] + nums[j]
				wanted := -cur
				set[wanted] = []int{nums[i], nums[j]}
			}
		}
	}
	return ans
}

func clear(set map[int][]int) {
	for k := range set {
		delete(set, k)
	}
}
