package valid

import "github.com/catorpilor/leetcode/utils"

func isValid(root *utils.TreeNode, arr []int) bool {
    return useDFS(root, arr, 0, len(arr))
}

// useDFS time complexity O(N), space complexity O(lgN)
func useDFS(root *utils.TreeNode, arr []int, pos, l int) bool {
    if root == nil {
        return l == 0
    }
    if pos == l-1 && root.Left == nil && root.Right == nil && root.Val == arr[pos] {
        return true
    }
    if pos < l-1 {
        return useDFS(root.Left, arr, pos+1, l) || useDFS(root.Right, arr, pos+1, l)
    }
    return false
}
