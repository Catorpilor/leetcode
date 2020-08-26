package post

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
	"github.com/google/go-cmp/cmp"
)

func TestPostorderTraverse(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  []int
	}{
		{"nil root", nil, nil},
		{"single node", utils.ConstructTree([]int{1}), []int{1}},
		{"testcase1", utils.ConstructTree([]int{1, math.MinInt32, 2, 3}), []int{3, 2, 1}},
		{"testcase2", utils.ConstructTree([]int{1, 4, 3, 2}), []int{2, 4, 3, 1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := postOrder(tt.root)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
