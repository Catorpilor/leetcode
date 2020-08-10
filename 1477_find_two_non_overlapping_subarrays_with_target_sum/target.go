package target

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func minLength(arr []int, target int) int {
	return useOnepass(arr, target)
}

func useTwoPointers(arr []int, target int) int {
	return -1
}

// useOnepass time complexity O(N), space complexity O(N)
func useOnepass(arr []int, target int) int {
	n := len(arr)
	// best[i] represents the minimum subarray with target sum equal to target ending at postion i
	best := make([]int, n)
	for i := range best {
		best[i] = math.MaxInt32
	}
	var sum, start int
	ans, bestSofar := math.MaxInt32, math.MaxInt32
	for i := range arr {
		sum += arr[i]
		for sum > target {
			// resize the window
			sum -= arr[start]
			start++
		}
		if sum == target {
			if start > 0 && best[start-1] != math.MaxInt32 {
				ans = utils.Min(ans, best[start-1]+i-start+1)
			}
			bestSofar = utils.Min(bestSofar, i-start+1)
		}
		best[i] = bestSofar
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// useHashmap time complexity O(N), space complexity O(N)
func useHashmap(arr []int, target int) int {
	n := len(arr)
	set := make(map[int]int, n)
	set[0] = -1
	var sum int
	for i := range arr {
		sum += arr[i]
		set[sum] = i
	}
	sum = 0
	lsize, ans := math.MaxInt32, math.MaxInt32
	for i := 0; i < n; i++ {
		sum += arr[i]
		// there is a subarray ends i with sum = target
		if v, exists := set[sum-target]; exists {
			lsize = utils.Min(lsize, i-v)
		}
		// looking for another subarray with sum equals target starts with i+1
		if v, exists := set[sum+target]; exists {
			ans = utils.Min(ans, lsize+v-i)
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
