package bst

import (
    "testing"

    "github.com/catorpilor/leetcode/utils"
)

func TestSearch(t *testing.T) {
    st := []struct {
        name string
        root *utils.TreeNode
        val  int
        exp  *utils.TreeNode
    }{
        {"empty tree", nil, 3, nil},
        {"single node not found", utils.ConstructTree([]int{3}), 2, nil},
        {"val equal to the root", utils.ConstructTree([]int{3, 2, 4}), 3, utils.ConstructTree([]int{3, 2, 4})},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := search(tt.root, tt.val)
            if !utils.IsEqual(tt.exp, out) {
                t.Fatalf("with input tree: %v and val: %d wanted %v but got %v", utils.LevelOrderTravesal(tt.root),
                    tt.val, utils.LevelOrderTravesal(tt.exp), utils.LevelOrderTravesal(out))
            }
            t.Log("pass")
        })
    }
}
