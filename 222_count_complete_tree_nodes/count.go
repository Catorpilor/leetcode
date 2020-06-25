package count

import "github.com/catorpilor/LeetCode/utils"

func countNodes(root *utils.TreeNode) int {
	return useDfs(root)
}

// useDfs time complexity O(N), space complexity O(log(N))
func useDfs(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	lc := useDfs(root.Left)
	rc := useDfs(root.Right)
	return 1 + lc + rc
}

// useDfsWithCBSProp time compleixty O(log(N)*log(N)), space complexity O(logN)
func useDfsWithCBSProp(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	lr, rr := root, root
	var lh, rh int
	for lr != nil {
		lr = lr.Left
		lh++
	}
	for rr != nil {
		rr = rr.Right
		rh++
	}
	if lh == rh {
		return 1<<lh - 1
	}
	return 1 + useDfsWithCBSProp(root.Left) + useDfsWithCBSProp(root.Right)
}
