package vot

import (
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func verticalTraversal(root *utils.TreeNode) [][]int {
	return useBruteForce(root)
}

type pos struct {
	row, col int
}

// useBruteForce time complexity O(N), space complexity O(N)
func useBruteForce(root *utils.TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	levels := []*utils.TreeNode{root}
	set := make(map[*utils.TreeNode]pos)
	set[root] = pos{0, 0}
	finals := make(map[int][]*utils.TreeNode)
	finals[0] = append(finals[0], root)
	for len(levels) > 0 {
		var locals []*utils.TreeNode
		for _, node := range levels {
			pp := set[node]
			if node.Left != nil {
				set[node.Left] = pos{pp.row + 1, pp.col - 1}
				locals = append(locals, node.Left)
				finals[pp.col-1] = append(finals[pp.col-1], node.Left)
			}
			if node.Right != nil {
				set[node.Right] = pos{pp.row + 1, pp.col + 1}
				locals = append(locals, node.Right)
				finals[pp.col+1] = append(finals[pp.col+1], node.Right)
			}
		}
		levels = locals
	}
	keys := make([]int, 0, len(finals))
	for k := range finals {
		keys = append(keys, k)
		if len(finals[k]) > 0 {
			// sort nodes based on levels, if they are on the same level just sort by values.
			sort.Slice(finals[k], func(i, j int) bool {
				pi, pj := set[finals[k][i]], set[finals[k][j]]
				if pi.row < pj.row || pi.row == pj.row && finals[k][i].Val < finals[k][j].Val {
					return true
				}
				return false
			})
		}
	}
	sort.Ints(keys)
	ans = make([][]int, 0, len(finals))
	for _, k := range keys {
		tmp := make([]int, 0, len(finals[k]))
		for i := range finals[k] {
			tmp = append(tmp, finals[k][i].Val)
		}
		ans = append(ans, tmp)
	}
	return ans
}
