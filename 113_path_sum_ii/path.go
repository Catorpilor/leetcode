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

func useDFS(root *utils.TreeNode, sum int) [][]int {
	var res [][]int
	helper2(&res, root, sum, 0, []int{})
	return res
}

func helper2(res *[][]int, root *utils.TreeNode, sum, cur int, path []int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val+cur == sum {
			local := make([]int, len(path))
			copy(local, path)
			*res = append((*res), local)
			return
		}
	} else {
		helper2(res, root.Left, sum, cur+root.Val, path)
		helper2(res, root.Right, sum, cur+root.Val, path)
	}
	return
}
