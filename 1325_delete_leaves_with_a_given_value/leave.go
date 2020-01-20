package leave

import "github.com/catorpilor/leetcode/utils"

func removeLeafNodes(root *utils.TreeNode, target int) *utils.TreeNode {
    return recur(root, target)
}

// recur time comlexity O(N) N nodes, space complexity O(lgn)
func recur(root *utils.TreeNode, target int) *utils.TreeNode {
    if root == nil {
        return root
    }
    root.Left = recur(root.Left, target)
    root.Right = recur(root.Right, target)
    if root.Left == nil && root.Right == nil && root.Val == target {
        root = nil
    }
    return root
}
