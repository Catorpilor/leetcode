package symtree

import "github.com/catorpilor/leetcode/utils"

func isSymTree(root *utils.TreeNode) bool {
	// same as check tree are equal
	return dfs(root, root)
}

// dfs time complexity O(N), space complexity O(N)
func dfs(p, q *utils.TreeNode) bool {
	if (p == nil && q == nil) || (p != nil && q != nil && p.Val == q.Val && dfs(p.Left, q.Right) && dfs(p.Right, q.Left)) {
		return true
	}
	return false
}
