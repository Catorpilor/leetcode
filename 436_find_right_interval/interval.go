package interval

import (
	"fmt"
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func rightInterval(intervals [][]int) []int {
	return useBinarySearch(intervals)
}

func useBinarySearch(intervals [][]int) []int {
	n := len(intervals)
	if n < 1 {
		return nil
	}
	if n == 1 {
		return []int{-1}
	}
	ans := make([]int, n)
	begs := make([]int, 0, n)
	pos := make(map[int]int, n)
	for i, r := range intervals {
		pos[r[0]] = i
		begs = append(begs, r[0])
	}
	sort.Ints(begs)
	fmt.Println(begs, pos)
	var res int
	for _, r := range intervals {
		idx := utils.LowerBound(begs, r[1])
		if idx == n {
			res = -1
		} else {
			res = pos[begs[idx]]
		}
		ans[pos[r[0]]] = res
	}
	return ans
}
