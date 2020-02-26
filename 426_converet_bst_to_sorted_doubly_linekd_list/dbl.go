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
