package preorder

import "github.com/catorpilor/leetcode/utils"

func preorderTraversal(root *utils.TreeNode) []int {
	return useStack(root)
}

// useStack time complexity O(N), space complexity O(N)
func useStack(root *utils.TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	st := make([]*utils.TreeNode, 0, 1000)
	st = append(st, root)
	n := len(st)
	for n != 0 {
		head := st[n-1]
		st = st[:n-1]
		ans = append(ans, head.Val)
		if head.Right != nil {
			st = append(st, head.Right)
		}
		if head.Left != nil {
			st = append(st, head.Left)
		}
		n = len(st)
	}
	return ans
}

// useMorris time complexity O(N), space complexity O(1)
func useMorris(root *utils.TreeNode) []int {
	var ans []int
	cur := root
	for cur != nil {
		if cur.Left == nil {
			ans = append(ans, cur.Val)
			cur = cur.Right
		} else {
			pre := cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = cur
				ans = append(ans, cur.Val)
				cur = cur.Left
			} else {
				pre.Right = nil
				cur = cur.Right
			}
		}
	}
	return ans
}
