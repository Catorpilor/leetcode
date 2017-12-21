package path

import (
	"github.com/catorpilor/leetcode/utils"
)

func PathSum(root *utils.TreeNode, sum int) [][]int {
	// backtracking and dfs
	var ret [][]int
	if root == nil {
		return ret
	}
	var temp []int
	helper(root, sum, &temp, &ret)
	return ret
}

func helper(node *utils.TreeNode, k int, nodes *[]int, res *[][]int) {
	n := len(*nodes)
	*nodes = append(*nodes, node.Val)
	if node.Left == nil && node.Right == nil && k == node.Val {
		// find one
		bak := make([]int, len(*nodes))
		copy(bak, *nodes)
		*res = append(*res, bak)
	} else {
		if node.Left != nil {
			helper(node.Left, k-node.Val, nodes, res)
		}
		if node.Right != nil {
			helper(node.Right, k-node.Val, nodes, res)
		}
	}
	*nodes = (*nodes)[:n]
}
