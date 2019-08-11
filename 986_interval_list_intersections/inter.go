package inter

import "github.com/catorpilor/leetcode/utils"

func IntervalIntersection(A [][]int, B [][]int) [][]int {
	lenA, lenB := len(A), len(B)
	var start, end int
	res := [][]int{}
	// res = append(res, []int{start, end})
	var i, j int
	for i < lenA && j < lenB {
		start = utils.Max(A[i][0], B[j][0])
		end = utils.Min(A[i][1], B[j][1])
		if start <= end {
			res = append(res, []int{start, end})
		}
		if B[j][1] < A[i][1] {
			j++
		} else {
			i++
		}
	}
	return res
}
