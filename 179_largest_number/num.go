package num

import (
	"bytes"
	"sort"
	"strconv"
)

func largestNum(nums []int) string {
	return useBucket(nums)
}

// useBucket time complexity O(N+k*logk), space complexity O(N)
func useBucket(nums []int) string {
	bkt := [10][]int{}
	for _, num := range nums {
		tmp := num
		for tmp >= 10 {
			tmp /= 10
		}
		bkt[tmp] = append(bkt[tmp], num)
	}
	var buf bytes.Buffer
	for i := 9; i >= 0; i-- {
		// sort bkt[i] based on the combinantion of string values
		vec := bkt[i]
		if len(vec) > 1 {
			sort.Slice(vec, func(i, j int) bool {
				si, sj := strconv.Itoa(vec[i]), strconv.Itoa(vec[j])
				return si+sj > sj+si
			})
		}
		for _, num := range vec {
			buf.WriteString(strconv.Itoa(num))
		}
	}
	ans := buf.String()
	if ans[0] == '0' {
		return "0"
	}
	return ans
}

// useSorting time complexity O(N*K*logK) space complexity O(N) where K is the average length of each number.
func useSorting(nums []int) string {
	// just sort the nums based on the combination of two
	sort.Slice(nums, func(i, j int) bool {
		si, sj := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return si+sj > sj+si
	})
	var sb bytes.Buffer
	for _, num := range nums {
		sb.WriteString(strconv.Itoa(num))
	}
	ans := sb.String()
	if ans[0] == '0' {
		return "0"
	}
	return ans
}
