package wk

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	return usePrefixSum(mat, k)
}

// usePrefixSum time complexity O(MN), space complexity O(N)
func usePrefixSum(mat [][]int, k int) []int {
	m := len(mat)
	if m < 1 {
		return nil
	}
	n := len(mat[0])
	if n < 1 {
		return nil
	}
	type obj struct {
		idx, val int
	}
	psum := make([]obj, n)
	for i := 0; i < m; i++ {
		psum[i].idx = i
		for j := 0; j < n; j++ {
			psum[i].val += mat[i][j]
		}
	}
	// sort psum based on rules
	sort.Slice(psum, func(i, j int) bool {
		if psum[i].val == psum[j].val {
			return psum[i].idx < psum[j].idx
		}
		return psum[i].val < psum[j].val
	})
	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, psum[i].idx)
	}
	return ans
}
