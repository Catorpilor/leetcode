package bst

import (
    "math"
    "testing"

    "github.com/catorpilor/leetcode/utils"
)

func TestInsert(t *testing.T) {
    st := []struct {
        name string
        root *utils.TreeNode
        val  int
        exp  *utils.TreeNode
    }{
        {"empty tree", nil, 3, utils.ConstructTree([]int{3})},
        {"single node", utils.ConstructTree([]int{1}), 2, utils.ConstructTree([]int{1, math.MinInt32, 2})},
        {"testcase1", utils.ConstructTree([]int{4, 2, 7, 1, 3}), 5, utils.ConstructTree([]int{4, 2, 7, 1, 3, 5})},
    }

    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := insertTree(tt.root, tt.val)
            if !utils.IsEqual(tt.exp, out) {
                t.Fatalf("with root:%v and val:%d wanted %v but got %v",
                    utils.LevelOrderTravesal(tt.root), tt.val, utils.LevelOrderTravesal(tt.exp),
                    utils.LevelOrderTravesal(out))
            }
        })
    }
}

func TestInsertIter(t *testing.T) {
    st := []struct {
        name string
        root *utils.TreeNode
        val  int
        exp  *utils.TreeNode
    }{
        {"empty tree", nil, 3, utils.ConstructTree([]int{3})},
        {"single node", utils.ConstructTree([]int{1}), 2, utils.ConstructTree([]int{1, math.MinInt32, 2})},
        {"testcase1", utils.ConstructTree([]int{4, 2, 7, 1, 3}), 5, utils.ConstructTree([]int{4, 2, 7, 1, 3, 5})},
    }

    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            out := insertTreeIter(tt.root, tt.val)
            if !utils.IsEqual(tt.exp, out) {
                t.Fatalf("with root:%v and val:%d wanted %v but got %v",
                    utils.LevelOrderTravesal(tt.root), tt.val, utils.LevelOrderTravesal(tt.exp),
                    utils.LevelOrderTravesal(out))
            }
        })
    }
}
