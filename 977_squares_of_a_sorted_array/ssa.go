package ssa

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func sortedSquares(A []int) []int {
	// return []int{}
	// return bf(A)
	return twoPts(A)
}

func bf(A []int) []int {
	// Time Complexity O(n + nlg(n))
	// Space complexity O(n)
	res := make([]int, 0, len(A))
	for i := range A {
		res = append(res, A[i]*A[i])
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

func twoPts(A []int) []int {
	// time complexity O(N)
	// space complexity O(N)
	n := len(A)
	res := make([]int, n)
	l, r := 0, n-1
	for j := n - 1; j >= 0; j-- {
		if utils.Abs(A[l]) > utils.Abs(A[r]) {
			res[j] = A[l] * A[l]
			l++
		} else {
			res[j] = A[r] * A[r]
			r--
		}
	}
	return res
}
