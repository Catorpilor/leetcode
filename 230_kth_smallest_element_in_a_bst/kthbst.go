package kthbst

import (
	"github.com/catorpilor/leetcode/utils"
)

func kthSmallest(root *utils.TreeNode, k int) int {
	ret := utils.InorderTraversal(root)
	if len(ret) < k || k < 1 {
		return 0
	}
	return ret[k-1]
}

func kthSmallestIteration(root *utils.TreeNode, k int) int {
	if k < 1 || root == nil {
		return 0
	}
	st := utils.NewStack()
	cur := root
	for {
		for cur != nil {
			st.Push(cur)
			cur = cur.Left
		}
		cur = st.Pop().(*utils.TreeNode)
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
}
