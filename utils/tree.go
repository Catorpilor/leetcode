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

// InorderTraversal uses Morris Traversal Algorithm
// time complexity is O(n) and space complexity is O(1)
func InorderTraversal(root *TreeNode) []int {
	var ret []int
	cur := root
	for cur != nil {
		if cur.Left == nil {
			ret = append(ret, cur.Val)
			cur = cur.Right
		} else {
			predecessor := cur.Left
			for predecessor.Right != cur && predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = cur
				cur = cur.Left
			} else {
				predecessor.Right = nil
				ret = append(ret, cur.Val)
				cur = cur.Right
			}
		}
	}
	return ret
}

// PreorderTraversal uses Morris Traversal Algorithm
// time complexity is O(n) and space complexity is O(1)
func PreorderTraversal(root *TreeNode) []int {
	var ret []int
	cur := root
	for cur != nil {
		if cur.Left == nil {
			ret = append(ret, cur.Val)
			cur = cur.Right
		} else {
			predecessor := cur.Left
			for predecessor.Right != cur && predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = cur
				ret = append(ret, cur.Val)
				cur = cur.Left
			} else {
				predecessor.Right = nil
				cur = cur.Right
			}
		}
	}
	return ret
}

// PostorderTraversal uses Morris Traversal Algorithm
// time complexity is O(n) and space complexity is O(1)
func PostorderTraversal(root *TreeNode) []int {
	var ret []int
	dummy := &TreeNode{Left: root}
	var predecessor, first, middle, last *TreeNode
	cur := dummy
	for cur != nil {
		if cur.Left == nil {
			// ret = append(ret, cur.Val)
			cur = cur.Right
		} else {
			predecessor = cur.Left
			for predecessor.Right != cur && predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = cur
				// ret = append(ret, cur.Val)
				cur = cur.Left
			} else {
				// predeccessor found second time
				// reverse the right references in chain from predecessor to p
				first, middle = cur, cur.Left
				for middle != cur {
					last = middle.Right
					middle.Right = first
					first = middle
					middle = last
				}

				// visit the nodes from pred to p
				// again reverse the right references from pred to p
				first, middle = cur, predecessor
				for middle != cur {
					ret = append(ret, middle.Val)
					last = middle.Right
					middle.Right = first
					first = middle
					middle = last
				}
				predecessor.Right = nil
				cur = cur.Right
			}
		}
	}
	return ret
}

func IsEqual(l, r *TreeNode) bool {
	if (l == nil && r != nil) || (l != nil && r == nil) {
		return false
	}
	if l == nil && r == nil {
		return true
	}
	if l.Val != r.Val {
		return false
	}
	return IsEqual(l.Left, r.Left) && IsEqual(l.Right, r.Right)
}
