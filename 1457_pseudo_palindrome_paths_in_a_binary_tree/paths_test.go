package path

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestMaxNumberOfPalindromPath(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  int
	}{
		{"nil root", nil, 0},
		{"single node", utils.ConstructTree([]int{1}), 1},
		{"testcase1", utils.ConstructTree([]int{1, 2, 3}), 0},
		{"testcase2", utils.ConstructTree([]int{2, 3, 1, 3, 1, math.MinInt32, 1}), 2},
		{"testcase3", utils.ConstructTree([]int{2, 1, 1, 1, 3, math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32, 1}), 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxNumOfPalindromePath(tt.root)
			if out != tt.exp {
				t.Fatalf("with input root:%v wanted %d but got %d", utils.LevelOrderTravesal(tt.root), tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
