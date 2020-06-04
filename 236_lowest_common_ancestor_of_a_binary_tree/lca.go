package lca

import "github.com/catorpilor/leetcode/utils"

func lowsetCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	return useDfs(root, p, q)
}

// useDfs time complexity O(N), space complexity O(lgN)
func useDfs(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	leftCommon := useDfs(root.Left, p, q)
	rightCommon := useDfs(root.Right, p, q)
	if leftCommon != nil && rightCommon != nil {
		return root
	}
	if leftCommon != nil {
		return leftCommon
	}
	return rightCommon
}
