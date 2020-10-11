package bst

import "github.com/catorpilor/leetcode/utils"

func insertIntoBST(root *utils.TreeNode, val int) *utils.TreeNode {
	return helper(root, val)
}

// helper is a recursive approach. tiem complexity O(N), space complexity O(lgN)
func helper(root *utils.TreeNode, val int) *utils.TreeNode {
	if root == nil {
		return &utils.TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = helper(root.Right, val)
	} else {
		root.Left = helper(root.Left, val)
	}
	return root
}

// useIter time complexity O(N), space complexity O(1)
func useIter(root *utils.TreeNode, val int) *utils.TreeNode {
	cur := root
	if cur == nil {
		return &utils.TreeNode{Val: val}
	}
	for {
		if val > cur.Val {
			if cur.Right == nil {
				cur.Right = &utils.TreeNode{Val: val}
				break
			} else {
				cur = cur.Right
			}
		} else {
			if cur.Left == nil {
				cur.Left = &utils.TreeNode{Val: val}
				break
			} else {
				cur = cur.Left
			}
		}
	}
	return root
}
