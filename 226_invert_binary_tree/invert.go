package invert

import "github.com/catorpilor/leetcode/utils"

func InvertRecur(root *utils.TreeNode) *utils.TreeNode {
    if root == nil {
        return root
    }
    left, right := root.Left, root.Right
    root.Left = InvertRecur(right)
    root.Right = InvertRecur(left)
    return root
}

func Invert(root *utils.TreeNode) *utils.TreeNode {
    if root == nil {
        return root
    }
    s := utils.NewStack()
    s.Push(root)
    for !s.IsEmpty() {
        n := s.Pop().(*utils.TreeNode)
        left := n.Left
        n.Left = n.Right
        n.Right = left
        if n.Left != nil {
            s.Push(n.Left)
        }
        if n.Right != nil {
            s.Push(n.Right)
        }
    }
    return root
}
