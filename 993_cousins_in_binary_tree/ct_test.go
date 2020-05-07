package ct

import (
    "math"
    "testing"

    "github.com/catorpilor/leetcode/utils"
)

func TestIsCousins(t *testing.T) {
    st := []struct {
        name string
        root *utils.TreeNode
        x, y int
        exp  bool
    }{
        {"empty tree", nil, 1, 2, false},
        {"single node", utils.ConstructTree([]int{1}), 2, 3, false},
        {"testcase1", utils.ConstructTree([]int{1, 2, 3, 4}), 3, 4, false},
        {"testcase2", utils.ConstructTree([]int{1, 2, 3, math.MinInt32, 4, math.MinInt32, 5}), 5, 4, true},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := isCousins(tt.root, tt.x, tt.y)
            if out != tt.exp {
                t.Fatalf("with tree: %v and x:%d, y:%d wanted %t but got %t", utils.LevelOrderTravesal(tt.root), tt.x, tt.y, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
