package post

import "github.com/catorpilor/leetcode/utils"

func postOrder(root *utils.TreeNode) []int {
	return useRecursion(root)
}

// useRecursion time complexity O(N), space complexity O(lgN)
func useRecursion(root *utils.TreeNode) []int {
	var ans []int
	helper(root, &ans)
	return ans
}

func helper(root *utils.TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, ans)
	helper(root.Right, ans)
	*ans = append((*ans), root.Val)
}

// useStack time complexity O(N), space complexity O(N)
func useStack(root *utils.TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	st := make([]*utils.TreeNode, 0, 1000)
	st = append(st, root, root)
	n := len(st)
	for n != 0 {
		top := st[n-1]
		st = st[:n-1]
		n--
		if n != 0 && st[n-1] == top {
			if top.Right != nil {
				st = append(st, top.Right, top.Right)
			}
			if top.Left != nil {
				st = append(st, top.Left, top.Left)
			}
		} else {
			ans = append(ans, top.Val)
		}
		n = len(st)
	}
	return ans
}
