package sum

import (
	"github.com/catorpilor/leetcode/utils"
)

func SumNumbers(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root, 0)

}

func helper(node *utils.TreeNode, val int) int {
	if node == nil {
		return val
	}
	val = val*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return val
	}
	leftSum, rightSum := 0, 0
	if node.Left != nil {
		leftSum = helper(node.Left, val)
	}
	if node.Right != nil {
		rightSum = helper(node.Right, val)
	}
	return leftSum + rightSum
}
