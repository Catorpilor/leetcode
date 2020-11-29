package bst

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestFindTilt(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  int
	}{
		{"empty tree", nil, 0},
		{"single node", utils.ConstructTree([]int{1}), 0},
		{"testcase1", utils.ConstructTree([]int{1, 2, 3}), 1},
		{"testcase2", utils.ConstructTree([]int{4, 2, 9, 3, 5, math.MinInt32, 7}), 15},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findTilt(tt.root)
			if out != tt.exp {
				t.Fatalf("with input tree wanted %d  but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
