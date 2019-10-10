package kthbst

import (
    "math"
    "testing"

    "github.com/catorpilor/leetcode/utils"
)

func TestKthSmallest(t *testing.T) {
    st := []struct {
        name string
        root *utils.TreeNode
        k    int
        exp  int
    }{
        {"empty tree", nil, 1, 0},
        {"single node", utils.ConstructTree([]int{1}), 1, 1},
        {"testcase1", utils.ConstructTree([]int{3, 1, 4, math.MinInt32, 2}), 1, 1},
        {"testcase2", utils.ConstructTree([]int{5, 3, 6, 2, 4, math.MinInt32, math.MinInt32, 1}), 3, 3},
    }

    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := kthSmallest(tt.root, tt.k)
            if out != tt.exp {
                t.Fatalf("with input root %v and k %d wanted %d but got %d",
                    tt.root, tt.k, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
