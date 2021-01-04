package bbt

import (
    "math"
    "testing"

    "github.com/catorpilor/leetcode/utils"
)

func TestIsBalanced(t *testing.T) {
    st := []struct {
        name string
        root *utils.TreeNode
        exp  bool
    }{
        {"empty", nil, true},
        {"single node", utils.ConstructTree([]int{1}), true},
        {"testcase1", utils.ConstructTree([]int{1, 2, 2, 3, 3, math.MinInt32, math.MinInt32, 4, 4}), false},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := isBalancedBU(tt.root)
            if out != tt.exp {
                t.Fatalf("wanted %t but got %t", tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
