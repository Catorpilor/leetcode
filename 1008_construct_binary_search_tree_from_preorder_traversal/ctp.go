package ctp

import "github.com/catorpilor/leetcode/utils"

func constructTree(order []int) *utils.TreeNode {
    return useDFS(order)
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
