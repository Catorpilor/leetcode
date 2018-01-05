package search

import (
	"github.com/catorpilor/leetcode/utils"
)

func RecoverTree(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return utils.LevelOrderTravesal(root)
	}
	// inorder travesal
	// recursion so bset or normal space complex is O(lgN) and worst case is O(N)
	var first, second, prev *utils.TreeNode
	traverse(root, &first, &second, &prev)
	// fmt.Println(first, second)
	first.Val, second.Val = second.Val, first.Val
	return utils.LevelOrderTravesal(root)
}

func traverse(node *utils.TreeNode, first, second, prev **utils.TreeNode) {
	if node == nil {
		return
	}
	traverse(node.Left, first, second, prev)
	if *prev != nil && (*prev).Val >= node.Val {
		if *first == nil {
			*first = *prev
		}
		if *first != nil {
			*second = node
		}
	}
	*prev = node
	traverse(node.Right, first, second, prev)
}
