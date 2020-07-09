package tree

import (
	"github.com/catorpilor/leetcode/utils"
)

func widthOfBinaryTree(root *utils.TreeNode) int {
	return useLevelOrder(root)
}

// useLevelOrder time complexity O(N), space complexity O(N)
func useLevelOrder(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	q := utils.NewQueue()
	qIndex := utils.NewQueue()
	// q stores the node
	q.Enroll(root)
	// qIndex stores the index of node.
	// for a binary tree, it's left child index (2*parent), it's right child index (2*parent+1)
	qIndex.Enroll(1)
	for !q.IsEmpty() {
		count := q.Size()
		left, right := 0, 0
		for i := 0; i < count; i++ {
			node := q.Pull().(*utils.TreeNode)
			idx := qIndex.Pull().(int)
			if i == 0 {
				left = idx
			}
			if i == count-1 {
				right = idx
			}
			if node.Left != nil {
				q.Enroll(node.Left)
				qIndex.Enroll(2 * idx)
			}
			if node.Right != nil {
				q.Enroll(node.Right)
				qIndex.Enroll(2*idx + 1)
			}
		}
		ans = utils.Max(ans, right-left+1)
	}
	return ans
}

func height(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + utils.Max(height(root.Left), height(root.Right))
}
