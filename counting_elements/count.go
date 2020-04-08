package count

import "sort"

func countElements(arr []int) int {
	return useHashmap(arr)
}

// useHashmap time complexity O(N), space complexity O(1)
func useHashmap(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	set := make(map[int]bool, n)
	for _, num := range arr {
		set[num] = true
	}
	var ans int
	for _, num := range arr {
		if set[num+1] {
			ans++
		}
	}
	return ans
}

// useSort time complexity O(NlgN), space complexity O(1)
func useSort(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	sort.Ints(arr)
	prev, prevCount := -1, 0
	var ans int
	for _, num := range arr {
		if num == prev {
			prevCount++
		} else {
			if num == prev+1 {
				ans += prevCount
			}
			prev = num
			prevCount = 1
		}
	}
	return ans
}
