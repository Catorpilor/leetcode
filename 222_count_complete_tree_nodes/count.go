package count

import "github.com/catorpilor/LeetCode/utils"

func countNodes(root *utils.TreeNode) int {
	return useDfs(root)
}

// useDfs time complexity O(N), space complexity O(log(N))
func useDfs(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	lc := useDfs(root.Left)
	rc := useDfs(root.Right)
	return 1 + lc + rc
}
