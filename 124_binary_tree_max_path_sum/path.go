package path

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func MaxPathSum(root *utils.TreeNode) int {
	ret := math.MinInt32
	cnt := helper(root, &ret)
	return utils.Max(cnt, ret)
}

func helper(node *utils.TreeNode, curMax *int) int {
	// return the max sum  when node is the loweset ancester.
	if node == nil {
		return 0
	}
	// caculate the left and right sub-tree's max sum
	// for example [1,2,3,15,16]
	// node.Val -> 2
	left := utils.Max(0, helper(node.Left, curMax))   // left = 15, curMax = 15
	right := utils.Max(0, helper(node.Right, curMax)) // right = 16 curMax = 16
	// update curMax if there is a path
	*curMax = utils.Max(*curMax, left+right+node.Val) // curMax = 33 , path [2,15,16]
	return utils.Max(left, right) + node.Val          // return node's max sum 18
}
