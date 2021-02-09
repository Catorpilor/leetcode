package bst

import "github.com/catorpilor/leetcode/utils"

func convertBST(root *utils.TreeNode) *utils.TreeNode {
	var sum int
	useDFS(root, &sum)
	return root
}

// useDFS time complexity O(N), space complexity O(log(N))
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
