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

func RecoverTree2(root *utils.TreeNode) []int {
	// morris traversal
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var first, second, prev, cur *utils.TreeNode
	cur = root
	for cur != nil {
		if cur.Left == nil {
			if prev != nil && prev.Val >= cur.Val {
				if first == nil {
					first = prev
				}
				if first != nil {
					second = cur
				}
			}
			prev = cur
			cur = cur.Right
		} else {
			predecessor := cur.Left
			for predecessor.Right != cur && predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			// node := preInorder(cur)
			if predecessor.Right == nil {
				predecessor.Right = cur
				cur = cur.Left
			} else {
				predecessor.Right = nil
				if prev != nil && prev.Val >= cur.Val {
					if first == nil {
						first = prev
					}
					if first != nil {
						second = cur
					}
				}
				prev = cur
				cur = cur.Right
			}
		}
	}
	first.Val, second.Val = second.Val, first.Val
	return utils.LevelOrderTravesal(root)
}

// preInorder returns a previous n for node in the inorder traversal
func preInorder(node *utils.TreeNode) *utils.TreeNode {
	n := node.Left
	for n.Right != node && n.Right != nil {
		n = n.Right
	}
	return n
}
