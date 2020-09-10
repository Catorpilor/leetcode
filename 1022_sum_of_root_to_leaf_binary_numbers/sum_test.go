package sum

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestPathSum(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  int
	}{
		{"empty tree", nil, 0},
		{"single node", utils.ConstructTree([]int{1}), 1},
		{"testcase1", utils.ConstructTree([]int{1, 0, 0, 1, 0, 1, 1}), 19},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := pathSum(tt.root)
			if out != tt.exp {
				t.Fatalf("with input tree wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
