package splk

import "fmt"

func numsOfSubarray(nums []int, k int) int {
	return twoPointers(nums, k)
}

func twoPointers(nums []int, k int) int {
	n := len(nums)
	p := 1
	l, r := 0, 0
	var res int
	for r < n {
		p *= nums[r]
		for l <= r && p >= k {
			p /= nums[l]
			l++
		}
		res += r - l + 1
		r++
	}
	return res
}

func mine(nums []int, k int) int {
	n := len(nums)
	var res int
	if n < 1 {
		return res
	}
	// sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	opts := make([]int, n)
	for i := range opts {
		if nums[i] < k {
			opts[i] = 1
		} else {
			opts[i] = -1
		}
	}
	opts[0] = 0
	// fmt.Printf("before parse: opt: %v\n", opts)
	prodSofar := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > k {
			fmt.Printf("nums[%d] = %d, k: %d\n", i, nums[i], k)
			prodSofar = 1
			continue
		}
		temp := nums[i] * prodSofar
		if temp < k {
			if opts[i-1] != -1 {
				opts[i] = opts[i-1]
			}
			prodSofar = temp
			continue
		}
		fmt.Printf("current i:%d, temp: %d, k:%d \n", i, temp, k)
		prodSofar = nums[i]
		j := i - 1
		for j >= 0 {
			t := nums[j] * prodSofar
			if t >= k {
				break
			}
			prodSofar = t
			opts[i] = j
			j--
		}
	}
	// fmt.Printf("opts: %v\n", opts)
	for i := range opts {
		if opts[i] == -1 {
			continue
		}
		res += i - opts[i] + 1
	}
	return res
}
