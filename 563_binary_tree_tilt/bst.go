package bst

import "github.com/catorpilor/leetcode/utils"

func findTilt(root *utils.TreeNode) int {
	var ans int
	usePostOrder(root, &ans)
	return ans
}

// usePostOrder time complexity O(N), space complexity O(logN)
func usePostOrder(root *utils.TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := usePostOrder(root.Left, ans)
	right := usePostOrder(root.Right, ans)
	*ans += utils.Abs(left - right)
	return left + right + root.Val
}
