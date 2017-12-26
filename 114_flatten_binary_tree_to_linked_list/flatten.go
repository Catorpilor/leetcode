package flatten

import (
	"github.com/catorpilor/leetcode/utils"
)

func Flatten(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}
	q := utils.NewQueue()
	preOrder(root, q)
	_ = q.Pull()
	root.Left = nil
	prev := root
	for !q.IsEmpty() {
		n := q.Pull().(*utils.TreeNode)
		// fmt.Println(n.Val)
		prev.Right = n
		n.Left = nil
		prev = n
	}
	return utils.LevelOrderTravesal(root)

	// return ret
}

func preOrder(node *utils.TreeNode, q *utils.Queue) {
	if node == nil {
		return
	}
	q.Enroll(node)
	preOrder(node.Left, q)
	preOrder(node.Right, q)
}

func Flatten2(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}
	// bottom up recursion
	_ = reversedPreOrder(root, nil)
	return utils.LevelOrderTravesal(root)

}

func reversedPreOrder(node, prev *utils.TreeNode) *utils.TreeNode {
	if node == nil {
		return prev
	}
	prev = reversedPreOrder(node.Right, prev)
	prev = reversedPreOrder(node.Left, prev)
	node.Right = prev
	node.Left = nil
	prev = node
	return prev
}

func Flatten3(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		if cur.Left != nil {
			// find current node's preorder node that links to current node's right subtree
			prev := cur.Left
			for prev.Right != nil {
				prev = prev.Right
			}
			prev.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
	return utils.LevelOrderTravesal(root)
}
