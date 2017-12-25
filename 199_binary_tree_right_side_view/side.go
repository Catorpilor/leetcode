package side

import (
	"github.com/catorpilor/leetcode/utils"
)

func RightSideView(root *utils.TreeNode) []int {
	var ret []int
	var levels [][]int
	helper(root, &levels, 0)
	for i := range levels {
		n := len(levels[i])
		ret = append(ret, levels[i][n-1])
	}
	return ret
}

func helper(node *utils.TreeNode, ans *[][]int, level int) {
	if node == nil {
		return
	}
	if len(*ans) == level {
		*ans = append(*ans, []int{})
	}
	(*ans)[level] = append((*ans)[level], node.Val)
	helper(node.Left, ans, level+1)
	helper(node.Right, ans, level+1)
}

func RightSideView2(root *utils.TreeNode) []int {
	var ret []int
	var curLevel int
	helper2(root, &curLevel, 1, &ret)
	return ret
}

func helper2(node *utils.TreeNode, curLevel *int, level int, ans *[]int) {
	if node == nil {
		return
	}
	if *curLevel < level {
		*ans = append(*ans, node.Val)
		*curLevel = level
	}
	helper2(node.Right, curLevel, level+1, ans)
	helper2(node.Left, curLevel, level+1, ans)
}

func RightSideView3(root *utils.TreeNode) []int {
	var ret []int
	helper3(root, 0, &ret)
	return ret
}

func helper3(node *utils.TreeNode, level int, ans *[]int) {
	if node == nil {
		return
	}
	if len(*ans) == level {
		*ans = append(*ans, node.Val)
	}
	helper3(node.Right, level+1, ans)
	helper3(node.Left, level+1, ans)
}
