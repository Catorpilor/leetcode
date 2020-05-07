package ct

import "github.com/catorpilor/leetcode/utils"

func isCousins(root *utils.TreeNode, x, y int) bool {
    xd, xp := useDFS(root, nil, x, 0)
    yd, yp := useDFS(root, nil, y, 0)
    return xd == yd && xp != yp
}

// useDFS time complexity O(N), space complexity O(N)
func useDFS(root, parent *utils.TreeNode, x, dep int) (int, *utils.TreeNode) {
    if root == nil {
        // -1 means not fount
        return -1, nil
    }
    if root.Val == x {
        return dep, parent
    }
    ld, lp := useDFS(root.Left, root, x, dep+1)
    rd, rp := useDFS(root.Right, root, x, dep+1)
    if ld == -1 {
        return rd, rp
    }
    return ld, lp
}
