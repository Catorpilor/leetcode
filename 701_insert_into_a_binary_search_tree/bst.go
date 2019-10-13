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

func insertTreeIter(root *utils.TreeNode, val int) *utils.TreeNode {
    cur := root
    if cur == nil {
        return &utils.TreeNode{Val: val}
    }
    for {
        if val > cur.Val {
            if cur.Right == nil {
                cur.Right = &utils.TreeNode{Val: val}
                break
            } else {
                cur = cur.Right
            }
        } else {
            if cur.Left == nil {
                cur.Left = &utils.TreeNode{Val: val}
                break
            } else {
                cur = cur.Left
            }
        }
    }
    return root
}
