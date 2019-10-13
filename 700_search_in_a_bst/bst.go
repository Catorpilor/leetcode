package bst

import "github.com/catorpilor/leetcode/utils"

func search(root *utils.TreeNode, val int) *utils.TreeNode {
    cur := root
    if cur == nil || cur.Val == val {
        return cur
    }
    for cur != nil {
        if cur.Val > val {
            cur = cur.Left
        } else if cur.Val < val {
            cur = cur.Right
        } else {
            return cur
        }
    }
    return nil
}
