package bst

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestConvertBST(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  *utils.TreeNode
	}{
		{"nil tree", nil, nil},
		{"single node", utils.ConstructTree([]int{1}), utils.ConstructTree([]int{1})},
		{"testcase1", utils.ConstructTree([]int{4, 1, 6, 0, 2, 5, 7, math.MinInt32, math.MinInt32, math.MinInt32, 3, math.MinInt32, math.MinInt32, math.MinInt32, 8}), utils.ConstructTree([]int{30, 36, 21, 36, 25, 26, 15, math.MinInt32, math.MinInt32, math.MinInt32, 33, math.MinInt32, math.MinInt32, math.MinInt32, 8})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := convertBST(tt.root)
			if !utils.IsEqual(tt.exp, out) {
				t.Fatalf("wanted %v but got %v", utils.LevelOrderTravesal(tt.exp), utils.LevelOrderTravesal(out))
			}
			t.Log("pass")
		})
	}
}
