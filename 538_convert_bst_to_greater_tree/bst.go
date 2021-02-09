package bst

import "github.com/catorpilor/leetcode/utils"

func convertBST(root *utils.TreeNode) *utils.TreeNode {
	var sum int
	useDFS(root, &sum)
	return root
}

func useDFS(root *utils.TreeNode, gsum *int) {
	if root == nil {
		return
	}
	ori := root.Val
	useDFS(root.Right, gsum)
	root.Val += (*gsum)
	*gsum += ori
	useDFS(root.Left, gsum)
}
