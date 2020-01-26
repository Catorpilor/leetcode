package btzz

import (
	"math"
	"reflect"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestZigzagLevelOrder(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  [][]int
	}{
		{"empty tree", nil, [][]int{}},
		{"single node", utils.ConstructTree([]int{1}), [][]int{[]int{1}}},
		{"testcase1", utils.ConstructTree([]int{1, 2, 3}), [][]int{[]int{1}, []int{3, 2}}},
		{"testcase2", utils.ConstructTree([]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}), [][]int{[]int{3}, []int{20, 9}, []int{15, 7}}},
		{"failed1", utils.ConstructTree([]int{0, 2, 4, 1, math.MinInt32, 3, -1, 5, 1, math.MinInt32, 6, math.MinInt32, 8}), [][]int{[]int{0}, []int{4, 2}, []int{1, 3, -1}, []int{8, 6, 1, 5}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := zigzagLevel(tt.root)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with tree: %v wanted %v but got %v", utils.LevelOrderTravesal(tt.root), tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
