package btzz

import "github.com/catorpilor/leetcode/utils"

func zigzagLevel(root *utils.TreeNode) [][]int {
	return bfs(root)
}

func bfs(root *utils.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	res = append(res, []int{root.Val})
	leveld := make([]*utils.TreeNode, 0, 1000)
	leveld = append(leveld, root)
	zag := true
	for len(leveld) != 0 {
		next := make([]*utils.TreeNode, 0, 1000)
		for _, node := range leveld {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		lnodes := make([]int, len(next))
		for i, node := range next {
			lnodes[i] = node.Val
		}
		if zag {
			// swap vals
			for i, j := 0, len(next)-1; i < j; i, j = i+1, j-1 {
				lnodes[i], lnodes[j] = lnodes[j], lnodes[i]
			}
			zag = false
		}
		if len(lnodes) != 0 {
			res = append(res, lnodes)
		}
		leveld = next
	}
	return res
}
