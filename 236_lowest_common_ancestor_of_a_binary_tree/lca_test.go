package lca

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestLowestCommonAncestor(t *testing.T) {
	st := []struct {
		name       string
		root, p, q *utils.TreeNode
		exp        *utils.TreeNode
	}{
		{"empty tree", nil, &utils.TreeNode{Val: 1}, &utils.TreeNode{Val: 2}, nil},
		{"testcaes1", utils.ConstructTree([]int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}), &utils.TreeNode{Val: 5}, &utils.TreeNode{Val: 1}, &utils.TreeNode{Val: 3}},
		{"testcaes2", utils.ConstructTree([]int{3, 5, 1, 6, 2, 0, 8, math.MinInt32, math.MinInt32, 7, 4}), &utils.TreeNode{Val: 5}, &utils.TreeNode{Val: 4}, &utils.TreeNode{Val: 5}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := lowsetCommonAncestor(tt.root, tt.p, tt.q)
			if out != nil && tt.exp != nil && out.Val != tt.exp.Val {
				t.Fatalf("with input tree: %v, p:%v and q:%v wanted %v but got %v", utils.LevelOrderTravesal(tt.root), tt.p, tt.q, tt.exp, out)
			}
			t.Log("pass")
		})
	}

}
