package tree

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestWidthOfBS(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  int
	}{
		{"empty tree", nil, 0},
		{"single node", utils.ConstructTree([]int{1}), 1},
		{"testcase1", utils.ConstructTree([]int{1, 2, math.MinInt32, 3}), 1},
		{"testcase2", utils.ConstructTree([]int{1, 2, 3, 4}), 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := widthOfBinaryTree(tt.root)
			if out != tt.exp {
				t.Fatalf("with input tree: %v, wanted %d but got %d", utils.LevelOrderTravesal(tt.root), tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
