package btc

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestMinCameras(t *testing.T) {
	st := []struct {
		name  string
		nodes []int
		root  *utils.TreeNode
		exp   int
	}{
		{"empty tree", []int{}, nil, 0},
		{"single node", []int{1}, utils.ConstructTree([]int{1}), 1},
		{"testcase1", []int{0, 0, math.MinInt32, 0, 0}, utils.ConstructTree([]int{0, 0, math.MinInt32, 0, 0}), 1},
		{"testcase2", []int{0, 0, math.MinInt32, 0, math.MinInt32, 0, math.MinInt32, math.MinInt32, 0}, utils.ConstructTree([]int{0, 0, math.MinInt32, 0, math.MinInt32, 0, math.MinInt32, math.MinInt32, 0}), 2},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := MinCameras(tt.root)
			if out != tt.exp {
				t.Fatalf("with input node slice: %v wanted %d but got %d", tt.nodes, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
