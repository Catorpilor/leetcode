package kthbst

import "github.com/catorpilor/leetcode/utils"

func kthSmallest(root *utils.TreeNode, k int) int {
	ret := utils.InorderTraversal(root)
	if len(ret) < k || k < 1 {
		return 0
	}
	return ret[k-1]
}
