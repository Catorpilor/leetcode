package bst

import "github.com/catorpilor/leetcode/utils"

func insertTree(root *utils.TreeNode, val int) *utils.TreeNode {
    if root == nil {
        return &utils.TreeNode{Val: val}
    }
    if root.Val < val {
        root.Right = insertTree(root.Right, val)
    } else {
        root.Left = insertTree(root.Left, val)
    }
    return root
}
