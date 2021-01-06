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
        {"single node", utils.ConstructTree([]int{1}), utils.ConstructTree([]int{1})},
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
