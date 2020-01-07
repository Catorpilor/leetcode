package splk

import (
	"fmt"
	"sort"
)

func numsOfSubarray(nums []int, k int) int {
	n := len(nums)
	var res int
	if n < 1 {
		return res
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	opts := make([]int, n)
	for i := range opts {
		if nums[i] < k {
			opts[i] = i
		} else {
			opts[i] = -1
		}
	}
	prodSofar := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > k {
			break
		}
		temp := nums[i] * prodSofar
		if temp < k {
			opts[i] = opts[i-1]
			prodSofar = temp
			continue
		}
		prodSofar = nums[i]
		j := i - 1
		for j >= 0 {
			t := nums[j] * prodSofar
			if t >= k {
				break
			}
			prodSofar = t
			j--
		}
		if j != i-1 {
			opts[i] = j
		}
	}
	fmt.Printf("opts: %v\n", opts)
	for i := range opts {
		if opts[i] == -1 {
			break
		}
		res += i - opts[i] + 1
	}
	return res
}
