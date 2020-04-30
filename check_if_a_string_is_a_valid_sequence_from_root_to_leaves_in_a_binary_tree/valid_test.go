package valid

import (
    "math"
    "testing"

    "github.com/catorpilor/leetcode/utils"
)

func TestIsValidPath(t *testing.T) {
    st := []struct {
        name string
        root *utils.TreeNode
        arr  []int
        exp  bool
    }{
        {"empty tree", nil, []int{1, 2, 3}, false},
        {"empty arr", utils.ConstructTree([]int{1}), nil, false},
        {"both empty", nil, nil, true},
        {"testcase1", utils.ConstructTree([]int{8, 3, math.MinInt32, 2, 1, 5, 4}), []int{8}, false},
        {"testcase2", utils.ConstructTree([]int{1}), []int{1}, true},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := isValid(tt.root, tt.arr)
            if out != tt.exp {
                t.Fatalf("with root: %v and arr: %v wanted %t but got %t", utils.LevelOrderTravesal(tt.root), tt.arr, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}
