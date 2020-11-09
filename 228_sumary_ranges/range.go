package ranger

import (
	"fmt"
	"strconv"
)

func SummaryRanges(nums []int) []string {
	n := len(nums)
	// var ret []string
	if n < 1 {
		return []string{}
	}
	ret := make([]string, 0, n)
	var ps int
	i := 1
	for i < n {
		if nums[i] != nums[i-1]+1 {
			// a gap
			ret = append(ret, helper(nums, ps, i-1))
			ps = i
		}
		i += 1
	}
	ret = append(ret, helper(nums, ps, n-1))
	return ret
}

func helper(nums []int, start, end int) string {
	if start == end {
		return fmt.Sprintf("%d", nums[start])
	}
	return fmt.Sprintf("%d->%d", nums[start], nums[end])
}

// useIterator time complexity O(N), space complexity O(1)
func useIterator(nums []int) []string {
	n := len(nums)
	if n < 1 {
		return nil
	}
	res := make([]string, 0, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if nums[i]-nums[r] == 1 {
			r++
		} else {
			if l == r {
				res = append(res, strconv.Itoa(nums[l]))
				l++
				r++
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[l], nums[r]))
				r++
				l = r
			}
		}
	}
	if l == r {
		res = append(res, strconv.Itoa(nums[l]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[l], nums[r]))
	}
	return res
}

// useSimpleIter time complexity O(N), space complexity O(1)
func useSimpleIter(nums []int) []string {
	n := len(nums)
	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		prev := nums[i]
		for i+1 < n && nums[i+1]-nums[i] == 1 {
			i++
		}
		if prev != nums[i] {
			res = append(res, fmt.Sprintf("%d->%d", prev, nums[i]))
		} else {
			res = append(res, strconv.Itoa(nums[i]))
		}
	}
	return res
}
