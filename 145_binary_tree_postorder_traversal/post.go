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
