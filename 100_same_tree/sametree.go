package sametree

import "github.com/catorpilor/leetcode/utils"

func isSameTree(p, q *utils.TreeNode) bool {
	return dfs(p, q)
}

// dfs time complexity O(N), space complexity O(lgN)
func dfs(p, q *utils.TreeNode) bool {
	if (p == nil && q == nil) || (p != nil && q != nil && p.Val == q.Val && dfs(p.Left, q.Left) && dfs(p.Right, q.Right)) {
		return true
	}
	return false
}
