package vot

import (
	"fmt"
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func verticalTraversal(root *utils.TreeNode) [][]int {
	return useBruteForce(root)
}

type pos struct {
	row, col int
}

func useBruteForce(root *utils.TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	levels := []*utils.TreeNode{root}
	set := make(map[*utils.TreeNode]pos)
	set[root] = pos{0, 0}
	finals := make(map[int][]int)
	finals[0] = append(finals[0], root.Val)
	for len(levels) > 0 {
		var locals []*utils.TreeNode
		for _, node := range levels {
			pp := set[node]
			fmt.Printf("processing node:%d, pos:%v\n", node.Val, pp)
			if node.Left != nil {
				set[node.Left] = pos{pp.row + 1, pp.col - 1}
				locals = append(locals, node.Left)
				finals[pp.col-1] = append(finals[pp.col-1], node.Left.Val)
			}
			if node.Right != nil {
				set[node.Right] = pos{pp.row + 1, pp.col + 1}
				locals = append(locals, node.Right)
				finals[pp.col+1] = append(finals[pp.col+1], node.Right.Val)
			}
		}
		levels = locals
	}
	fmt.Printf("finals: %v\n", finals)
	keys := make([]int, 0, len(finals))
	for k := range finals {
		fmt.Printf("cur key in finals %d\n", k)
		keys = append(keys, k)
	}
	fmt.Printf("before sorting keys: %v\n", keys)
	sort.Ints(keys)
	fmt.Printf("sorted keys: %v\n", keys)
	ans = make([][]int, 0, len(finals))
	for _, k := range keys {
		ans = append(ans, finals[k])
		fmt.Printf("col: %d, ans: %v\n", k, ans)
	}
	return ans
}
