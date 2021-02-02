package trim

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestTrimBST(t *testing.T) {
	st := []struct {
		name      string
		root      *utils.TreeNode
		low, high int
		exp       *utils.TreeNode
	}{
		{"empty", nil, 1, 2, nil},
		{"testcase1", utils.ConstructTree([]int{1, 0, 2}), 1, 2, utils.ConstructTree([]int{2, math.MinInt32, 2})},
		{"testcaes2", utils.ConstructTree([]int{3, 0, 4, math.MinInt32, 2, math.MinInt32, math.MinInt32, math.MinInt32, 2, 1}), 1, 3,
			utils.ConstructTree([]int{3, 2, math.MinInt32, 1})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := trimBST(tt.root, tt.low, tt.high)
			if !utils.IsEqual(tt.exp, out) {
				t.Fatal("trimed error")
			}
			t.Log("pass")
		})
	}

}
