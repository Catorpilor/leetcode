package leaves

import (
	"github.com/catorpilor/leetcode/utils"
)

func FindLeaves(root *utils.TreeNode) [][]int {
	var res [][]int
	height(root, &res)
	return res
}

func height(root *utils.TreeNode, res *[][]int) int {
	if root == nil {
		return -1
	}
	level := 1 + utils.Max(height(root.Left, res), height(root.Right, res))
	if level == len(*res) {
		*res = append(*res, []int{})
	}
	*res[level] = append((*res)[level], root.Val)
	return level
}
