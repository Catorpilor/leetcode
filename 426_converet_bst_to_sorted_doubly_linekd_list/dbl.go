package dbl

import (
	"fmt"

	"github.com/catorpilor/leetcode/utils"
)

func treeToDoubly(root *utils.TreeNode) *utils.TreeNode {
	return useRecur(root)
}

// beg, end represents the minNode, maxNode so far
var beg, end *utils.TreeNode

func useRecur(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return root
	}
	inOrder(root)
	beg.Left = end
	end.Right = beg
	return beg
}

func inOrderTraversal(root *utils.TreeNode) {
	if root.Left != nil {
		inOrderTraversal(root.Left)
	}
	fmt.Println(root.Val)
	if root.Right != nil {
		inOrderTraversal(root.Right)
	}
}

// inOrder time complexity O(N), space complexity O(N)
func inOrder(cur *utils.TreeNode) {
	if cur != nil {
		inOrder(cur.Left)
		if end != nil {
			end.Right = cur
			cur.Left = end
		} else {
			beg = cur
		}
		end = cur
		inOrder(cur.Right)
	}
}

// useStack time complexity O(N), space complexity O(N)
func useStack(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return root
	}
	dummy := &utils.TreeNode{}
	st := make([]*utils.TreeNode, 0, 2000)
	prev := dummy
	cur := root
	n := len(st)
	for n > 0 || cur != nil {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		n = len(st)
		// pop from stack
		cur = st[n-1]
		st = st[:n-1]
		n--
		// update pointers
		prev.Right = cur
		cur.Left = prev
		prev = cur
		cur = cur.Right
	}
	// prev point to the last node, make it a cycle
	dummy.Right.Left = prev
	prev.Right = dummy.Right
	return dummy.Right
}
