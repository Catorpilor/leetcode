package utils

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// we use math.MinInt32 to represents null
func ConstructTree(s []int) *TreeNode {
	n := len(s)
	if n == 0 {
		return nil
	}
	q := NewQueue()
	var i int
	root := &TreeNode{Val: s[i]}
	q.Enroll(root)
	i++
	for !q.IsEmpty() && i < n {
		node := q.Pull().(*TreeNode)
		if s[i] != math.MinInt32 {
			node.Left = &TreeNode{Val: s[i]}
			q.Enroll(node.Left)
		}
		i++
		if i < n && s[i] != math.MinInt32 {
			node.Right = &TreeNode{Val: s[i]}
			q.Enroll(node.Right)
		}
		i++
	}
	return root
}

func LevelOrderTravesal(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	q := NewQueue()
	q.Enroll(root)
	for !q.IsEmpty() {
		v := q.Pull().(*TreeNode)
		ret = append(ret, v.Val)
		if v.Left != nil {
			q.Enroll(v.Left)
		}
		if v.Right != nil {
			q.Enroll(v.Right)
		}
	}
	return ret
}
