package sum

import "github.com/catorpilor/leetcode/utils"

func pathSum(root *utils.TreeNode, sum int) int {
	return useBruteForce(root, sum)
}

// useBruteForce space complexity O(N), time complexity O(N^2) worst case, and O(N*Log(N)) in best case.
func useBruteForce(root *utils.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helper(root, sum) + useBruteForce(root.Left, sum) + useBruteForce(root.Right, sum)
}

func helper(root *utils.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	extra := 0
	if root.Val == sum {
		extra = 1
	}
	return extra + helper(root.Left, sum-root.Val) + helper(root.Right, sum-root.Val)
}
