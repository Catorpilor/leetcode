package maxdiffs

import "github.com/catorpilor/leetcode/utils"

func maxDiff(root *utils.TreeNode) int {
    return useDfs(root)
}

// useDfs time complexity O(N), space complexity O(N)
func useDfs(root *utils.TreeNode) int {
    if root == nil {
        return 0
    }
    return helper(root, root.Val, root.Val)

}

func helper(root *utils.TreeNode, min, max int) int {
    if root == nil {
        return max - min
    }
    max = utils.Max(max, root.Val)
    min = utils.Min(min, root.Val)
    return utils.Max(helper(root.Left, min, max), helper(root.Right, min, max))
}
