package tree

import "github.com/catorpilor/leetcode/utils"

func increasingBST(root *utils.TreeNode) *utils.TreeNode {
    return useDFS(root)
}

// useDFS time complexity O(N), space complexity O(lgN)
func useDFS(root *utils.TreeNode) *utils.TreeNode {
    if root == nil || root.Left == nil && root.Right == nil {
        return root
    }
    var lh, lt *utils.TreeNode
    if root.Left != nil {
        lh, lt = helper(root.Left)
    }
    var rh *utils.TreeNode
    if root.Right != nil {
        rh, _ = helper(root.Right)
    }
    if lt != nil {
        lt.Right = root
    } else {
        lh = root
    }
    // rh, _ := helper(root.Right)
    root.Left = nil
    root.Right = rh
    return lh
}

func helper(root *utils.TreeNode) (head, tail *utils.TreeNode) {
    if root.Left == nil {
        head = root
        if root.Right != nil {
            lh, lt := helper(root.Right)
            root.Right = lh
            tail = lt
        } else {
            tail = root
        }
    } else {
        var lt *utils.TreeNode
        head, lt = helper(root.Left)
        lt.Right = root
        root.Left = nil
        if root.Right == nil {
            tail = root
        } else {
            rh, rt := helper(root.Right)
            root.Right = rh
            tail = rt
        }
    }
    return
}
