package bst

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func IsValidBST(root *utils.TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)

}

func helper(node *utils.TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
}
