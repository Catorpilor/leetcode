package level

import (
	"github.com/catorpilor/leetcode/utils"
)

func LevelOrder(root *utils.TreeNode) [][]int {
	var ret [][]int
	helper(root, &ret, 0)
	return ret
}

func helper(node *utils.TreeNode, ans *[][]int, level int) {
	if node == nil {
		return
	}
	if level == len(*ans) {
		*ans = append(*ans, []int{})
	}
	(*ans)[level] = append((*ans)[level], node.Val)
	helper(node.Left, ans, level+1)
	helper(node.Right, ans, level+1)
}
