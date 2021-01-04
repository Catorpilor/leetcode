package bbt

import "github.com/catorpilor/leetcode/utils"

// topdown approach
func isBalancedTD(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	l := depth(root.Left)
	r := depth(root.Right)

	return (l-r <= 1 && l-r >= -1) && isBalancedTD(root.Left) && isBalancedTD(root.Right)
}

func depth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	return utils.Max(l, r) + 1
}

// bottom up
// isBalancedBU time complexity O(N), space complexity O(lgN)
func isBalancedBU(root *utils.TreeNode) bool {
	return helper(root) != -1
}

func helper(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	l := helper(root.Left)
	if l == -1 {
		return -1
	}
	r := helper(root.Right)
	if r == -1 {
		return -1
	}
	if (l-r > 1) || (l-r < -1) {
		return -1
	}
	return utils.Max(l, r) + 1
}
