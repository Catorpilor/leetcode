package ranger

import "fmt"

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
