package kth

import (
	"math"
)

type TreeAncestor struct {
	// for now just use brute force
	n         int
	avgHeight int
	parent    []int
	dp        [][]int
}

func Construct(n int, parent []int) *TreeAncestor {
	h := int(math.Log2(float64(n)))
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, h+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
		dp[i][0] = parent[i]
	}
	for l := 1; l <= h; l++ {
		for i := 1; i < n; i++ {
			if dp[i][l-1] != -1 {
				dp[i][l] = dp[dp[i][l-1]][l-1]
			}
		}
	}

	return &TreeAncestor{
		n:         n,
		parent:    parent,
		avgHeight: h,
		dp:        dp,
	}
}

// useBinaryLifting time complexity O(logN)
func (this *TreeAncestor) useBinaryLifting(node, k int) int {
	cur := node
	for i := this.avgHeight; i >= 0; i-- {
		if this.dp[cur][i] != -1 && (1<<i) <= k {
			cur = this.dp[cur][i]
			k -= 1 << i
		}
	}
	if k > 0 {
		return -1
	}
	return cur
}

func (this *TreeAncestor) GetKthAncestor(node, k int) int {
	tmp := node
	for i := 1; i <= k; i++ {
		if this.parent[tmp] != -1 {
			tmp = this.parent[tmp]
		} else {
			return -1
		}
	}
	return tmp
}
