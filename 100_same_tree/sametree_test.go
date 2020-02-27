package sametree

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestIsSameTree(t *testing.T) {
	st := []struct {
		name string
		p, q *utils.TreeNode
		exp  bool
	}{
		{"both empty", nil, nil, true},
		{"one is empty", utils.ConstructTree([]int{1}), nil, false},
		{"equal", utils.ConstructTree([]int{4, 1, 7, 3, 2}), utils.ConstructTree([]int{4, 1, 7, 3, 2}), true},
		{"single node", utils.ConstructTree([]int{4}), utils.ConstructTree([]int{4}), true},
		{"testcase1", utils.ConstructTree([]int{4, 1, 7, 3, 2}), utils.ConstructTree([]int{4, 1, 7, 3, 2, 5}), false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isSameTree(tt.p, tt.q)
			if out != tt.exp {
				t.Fatalf("with tree p: %v and q: %v wanted %t but got %t", utils.InorderTraversal(tt.p),
					utils.InorderTraversal(tt.q), tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
