package vot

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
	"github.com/google/go-cmp/cmp"
)

func TestVerticalTraversal(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  [][]int
	}{
		{"nil", nil, nil},
		{"single node", utils.ConstructTree([]int{1}), [][]int{[]int{1}}},
		{"testcase1", utils.ConstructTree([]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}), [][]int{[]int{9}, []int{3, 15}, []int{20}, []int{7}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := verticalTraversal(tt.root)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(+wanted, -got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
