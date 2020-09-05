package all

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
	"github.com/google/go-cmp/cmp"
)

func TestAllElements(t *testing.T) {
	st := []struct {
		name         string
		root1, root2 *utils.TreeNode
		exp          []int
	}{
		{"all nil", nil, nil, nil},
		{"root1 is nil", nil, utils.ConstructTree([]int{2, 1, 3}), []int{1, 2, 3}},
		{"testcase1", utils.ConstructTree([]int{2, 1, 4}), utils.ConstructTree([]int{1, 0, 3}), []int{0, 1, 1, 2, 3, 4}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := allElements(tt.root1, tt.root2)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
