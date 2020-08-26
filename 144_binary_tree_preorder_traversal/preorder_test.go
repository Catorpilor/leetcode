package preorder

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
	"github.com/google/go-cmp/cmp"
)

func TestPreorderTraversal(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  []int
	}{
		{"nil tree", nil, nil},
		{"single node", utils.ConstructTree([]int{1}), []int{1}},
		{"testcase1", utils.ConstructTree([]int{1, math.MinInt32, 2, 3}), []int{1, 2, 3}},
		{"failed1", utils.ConstructTree([]int{1, 4, 3, 2}), []int{1, 4, 2, 3}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := preorderTraversal(tt.root)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
