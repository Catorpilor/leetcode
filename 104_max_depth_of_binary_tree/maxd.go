package maxd

import "github.com/catorpilor/leetcode/utils"

func maxDepth(root *utils.TreeNode) int {
	return useDfs(root)
}

// useDfs time complexity O(N), space complexity O(lgN)
func useDfs(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	l := useDfs(root.Left)
	r := useDfs(root.Right)
	return 1 + utils.Max(l, r)
}
