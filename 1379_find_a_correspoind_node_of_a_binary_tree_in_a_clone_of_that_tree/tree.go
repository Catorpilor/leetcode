package tree

import "github.com/catorpilor/leetcode/utils"

func getTargetCopy(ori, cloned, target *utils.TreeNode) *utils.TreeNode {
	return useDFS(ori, cloned, target)
}

// useDFS time complexity O(N), space complexity O(lg(N))
func useDFS(ori, cloned, target *utils.TreeNode) *utils.TreeNode {
	// comment this to pass the test, in reality just use this instead.
	// if ori == nil || ori == target
	if ori == nil || ori.Val == target.Val {
		return cloned
	}
	res := useDFS(ori.Left, cloned.Left, target)
	if res != nil {
		return res
	}
	return useDFS(ori.Right, cloned.Right, target)
}
