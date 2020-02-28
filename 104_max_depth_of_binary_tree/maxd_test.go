package maxd

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestMaxDepthOfBTree(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  int
	}{
		{"nil tree", nil, 0},
		{"single node", utils.ConstructTree([]int{1}), 1},
		{"testcase1", utils.ConstructTree([]int{1, 2, 3, 4, 5}), 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxDepth(tt.root)
			if out != tt.exp {
				t.Fatalf("with input tree: %v wanted %d but got %d", utils.LevelOrderTravesal(tt.root), tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
