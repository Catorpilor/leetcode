package seq

import "sort"

func LongestConsecutive(nums []int) int {
	// O(nlgn)
	n := len(nums)
	if n <= 1 {
		return n
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	ret, lenSoFar := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]+1 {
			lenSoFar++
			if lenSoFar > ret {
				ret = lenSoFar
			}
		} else if nums[i] == nums[i-1] {
			continue
		} else {

			lenSoFar = 1
		}
	}
	return ret

}

func LongestConsecutive2(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	// hm stores the sequence length for key v
	hm := make(map[int]int)
	var l, r int // l, r represents the left, right boundary of v-1, and v+1
	var ret, lenSofar int
	for _, v := range nums {
		if hm[v] == 0 {
			l, r = hm[v-1], hm[v+1]
			lenSofar = l + r + 1
			if lenSofar > ret {
				ret = lenSofar
			}
			hm[v] = lenSofar
			hm[v-l], hm[v+r] = lenSofar, lenSofar // update boudary
		}
	}
	return ret
}

func LongestConsecutive3(nums []int) int {
	// brute force
	n := len(nums)
	if n <= 1 {
		return n
	}
	var ret, lenSofar, curNum int

	for _, v := range nums {
		curNum, lenSofar = v, 1
		for helper(nums, curNum+1) {
			lenSofar, curNum = lenSofar+1, curNum+1
		}
		if lenSofar > ret {
			ret = lenSofar
		}
	}
	return ret
}
func helper(nums []int, v int) bool {
	for _, c := range nums {
		if c == v {
			return true
		}
	}
	return false
}
