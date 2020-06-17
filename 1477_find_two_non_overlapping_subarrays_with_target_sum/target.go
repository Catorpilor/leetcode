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

func useOnepass(arr []int, target int) int {
	n := len(arr)
	best := make([]int, n)
	for i := range best {
		best[i] = math.MaxInt32
	}
	var sum, start int
	ans, bestSofar := math.MaxInt32, math.MaxInt32
	for i := range arr {
		sum += arr[i]
		for sum > target {
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
