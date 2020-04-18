package maxdiffs

import (
    "math"
    "testing"

    "github.com/catorpilor/leetcode/utils"
)

func TestMaxDiff(t *testing.T) {
    st := []struct {
        name string
        root *utils.TreeNode
        exp  int
    }{
        {"empty tree", nil, 0},
        {"just root", utils.ConstructTree([]int{1}), 0},
        {"testcase1", utils.ConstructTree([]int{1, 2, 3, 4, 5}), 4},
        {"testcase2", utils.ConstructTree([]int{8, 3, 10, 1, 6, math.MinInt32, 14, math.MinInt32, math.MinInt32, 4, 7, 13}), 7},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := maxDiff(tt.root)
            if out != tt.exp {
                t.Fatalf("with input tree:%v, wanted %d but got %d", utils.LevelOrderTravesal(tt.root), tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
