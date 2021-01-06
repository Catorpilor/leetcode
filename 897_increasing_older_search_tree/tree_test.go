package tree

import (
    "math"
    "testing"

    "github.com/catorpilor/leetcode/utils"
)

func TestIncreasingBST(t *testing.T) {
    st := []struct {
        name string
        root *utils.TreeNode
        exp  *utils.TreeNode
    }{
        {"testcase1", utils.ConstructTree([]int{5, 1, 7}), utils.ConstructTree([]int{1, math.MinInt32, 5, math.MinInt32, 7})},
        {"empty tree", nil, nil},
        {"failed1", utils.ConstructTree([]int{106, math.MinInt32, 783}), utils.ConstructTree([]int{106, math.MinInt32, 783})},
        {"single node", utils.ConstructTree([]int{1}), utils.ConstructTree([]int{1})},
        {"testcase2", utils.ConstructTree([]int{5, 3, 6, 2, 4, math.MinInt32, 8, 1, math.MinInt32, math.MinInt32, math.MinInt32, 7, 9}),
            utils.ConstructTree([]int{1, math.MinInt32, 2, math.MinInt32, 3, math.MinInt32, 4, math.MinInt32, 5, math.MinInt32, 6, math.MinInt32, 7, math.MinInt32, 8, math.MinInt32, 9})},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := increasingBST(tt.root)
            if !utils.IsEqual(tt.exp, out) {
                t.Fatalf("wanted %v but got %v", utils.LevelOrderTravesal(tt.exp),
                    utils.LevelOrderTravesal(out))
            }
            t.Log("pass")
        })
    }
}
