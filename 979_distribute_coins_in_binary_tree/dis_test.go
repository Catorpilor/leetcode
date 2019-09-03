package dis

import (
	"testing"

	"github.com/catorpilor/LeetCode/utils"
)

func TestDistributeCoins(t *testing.T) {
	st := []struct {
		name string
		root *utils.TreeNode
		exp  int
	}{
		{"empty tree", nil, 0},
		{"single node", utils.ConstructTree([]int{1}), 0},
		{"testcase1", utils.ConstructTree([]int{3, 0, 0}), 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := DistributeCoins(tt.root)
			if out != tt.exp {
				t.Fatalf("with input tree %v, wanted %d but got %d", tt.root, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
