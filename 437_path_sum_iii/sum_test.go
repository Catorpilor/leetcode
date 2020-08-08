package sum

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestPathSum(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		sum  int
		exp  int
	}{
		{"nil tree", nil, 5, 0},
		{"testcase1", utils.ConstructTree([]int{10, 5, -3, 3, 2, math.MinInt32, 11, 3, -2, math.MinInt32, 1}), 8, 3},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := pathSum(tt.root, tt.sum)
			if out != tt.exp {
				t.Fatalf("with input tree wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
