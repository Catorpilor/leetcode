package dis

import "github.com/catorpilor/leetcode/utils"

// DistributeCoins distribute coins
func DistributeCoins(root *utils.TreeNode) int {
	var res int
	dfs(root, &res)
	return res
}

func dfs(root *utils.TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	// l ,r represents the total coins needs to be transfered for
	// root's substree
	l, r := dfs(root.Left, res), dfs(root.Right, res)
	*res += utils.Abs(l) + utils.Abs(r)
	return root.Val + l + r - 1
}
