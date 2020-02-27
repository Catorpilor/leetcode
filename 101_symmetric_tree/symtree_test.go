package symtree

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestIsSymTree(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  bool
	}{
		{"both empty", nil, true},
		{"one is empty", utils.ConstructTree([]int{1}), true},
		{"equal", utils.ConstructTree([]int{4, 1, 7, 3, 2}), false},
		{"testcase1", utils.ConstructTree([]int{4, 2, 2, 3, 4, 4, 3}), true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isSymTree(tt.root)
			if out != tt.exp {
				t.Fatalf("with input tree: %v wanted %t but got %t", utils.LevelOrderTravesal(tt.root), tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
