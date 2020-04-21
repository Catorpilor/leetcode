package ctp

import (
    "math"

    "github.com/catorpilor/leetcode/utils"
)

func constructTree(order []int) *utils.TreeNode {
    // return useDFS(order)
    var idx int
    return useDFSWithLinear(order, &idx, math.MaxInt32)
}

// useDFS time complexity O(N^2), space complexity O(N)
func useDFS(order []int) *utils.TreeNode {
    n := len(order)
    if n == 0 {
        return nil
    }
    // for example order=[1,2,3,4,5,6,7,....] it's unbalanced binary tree.
    root := &utils.TreeNode{Val: order[0]}
    pos := 1
    for pos < n {
        if order[pos] > order[0] {
            break
        }
        pos++
    }
    root.Left = useDFS(order[1:pos])
    root.Right = useDFS(order[pos:])
    return root
}

// useDFSWithLinear time complexity O(N), space complexity O(N)
func useDFSWithLinear(order []int, idx *int, max int) *utils.TreeNode {

    if (*idx) == len(order) || order[(*idx)] > max {
        return nil
    }
    root := &utils.TreeNode{Val: order[(*idx)]}
    (*idx)++
    root.Left = useDFSWithLinear(order, idx, root.Val)
    // idx need to be a reference, otherwise we're revisiting some indexes in the left subtree.
    root.Right = useDFSWithLinear(order, idx, max)
    return root
}
