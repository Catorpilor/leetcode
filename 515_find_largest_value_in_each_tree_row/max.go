package max

import (
	"github.com/catorpilor/leetcode/utils"
)

func LargestValue(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	helper(root, 0, &ret)
	return ret
}

func helper(node *utils.TreeNode, level int, ans *[]int) {
	if node == nil {
		return
	}
	if level == len(*ans) {
		*ans = append(*ans, node.Val)
	} else {
		(*ans)[level] = utils.Max((*ans)[level], node.Val)
	}
	if node.Left != nil {
		helper(node.Left, level+1, ans)
	}
	if node.Right != nil {
		helper(node.Right, level+1, ans)
	}
}
