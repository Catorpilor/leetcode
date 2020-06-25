package count

import (
	"testing"

	"github.com/catorpilor/LeetCode/utils"
)

func TestCountNodes(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  int
	}{
		{"emtpty tree", nil, 0},
		{"single node", utils.ConstructTree([]int{1}), 1},
		{"full complete tree", utils.ConstructTree([]int{1, 2, 3, 4, 5, 6, 7}), 7},
		{"testcase1", utils.ConstructTree([]int{1, 2, 3, 4, 5, 6}), 6},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := countNodes(tt.root)
			if out != tt.exp {
				t.Fatalf("with input tree wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestCountNodesWithCBSProp(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  int
	}{
		{"emtpty tree", nil, 0},
		{"single node", utils.ConstructTree([]int{1}), 1},
		{"full complete tree", utils.ConstructTree([]int{1, 2, 3, 4, 5, 6, 7}), 7},
		{"testcase1", utils.ConstructTree([]int{1, 2, 3, 4, 5, 6}), 6},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useDfsWithCBSProp(tt.root)
			if out != tt.exp {
				t.Fatalf("with input tree wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
