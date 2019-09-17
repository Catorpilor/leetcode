package ssa

import "sort"

func sortedSquares(A []int) []int {
	// return []int{}
	return bf(A)
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
