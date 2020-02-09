package leave

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestRemoveLeafNodes(t *testing.T) {
	st := []struct {
		name   string
		root   *utils.TreeNode
		target int
		exp    *utils.TreeNode
	}{
		{"empty tree", nil, 1, nil},
		{"one node equal target", utils.ConstructTree([]int{1}), 1, nil},
		{"root equals target", utils.ConstructTree([]int{1, 2, 3}), 1, utils.ConstructTree([]int{1, 2, 3})},
		{"testcase1", utils.ConstructTree([]int{1, 1, 1}), 1, nil},
		{"testcase2", utils.ConstructTree([]int{1, 2, 3, 2, math.MinInt32, 2, 4}), 2, utils.ConstructTree([]int{1, math.MinInt32,
			3, math.MinInt32, 4})},
	}
	for _, tt := range st {
		out := removeLeafNodes(tt.root, tt.target)
		if !utils.IsEqual(tt.exp, out) {
			t.Fatalf("with root: %v and target %d wanted %v but got %v", utils.LevelOrderTravesal(tt.root), tt.target,
				utils.LevelOrderTravesal(tt.exp), utils.LevelOrderTravesal(out))
		}
		t.Log("pass")
	}
}
