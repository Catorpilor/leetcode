package tree

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestBuildTree(t *testing.T) {
	st := []struct {
		name              string
		preorder, inorder []int
		exp               *utils.TreeNode
	}{
		{"both nil", nil, nil, nil},
		{"single node", []int{1}, []int{1}, utils.ConstructTree([]int{1})},
		{"testcase1", []int{1, 2, 4, 5, 3, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7}, utils.ConstructTree([]int{1, 2, 3, 4, 5, 6, 7})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := buildTree(tt.preorder, tt.inorder)
			if !utils.IsEqual(out, tt.exp) {
				t.Fatalf("with preorder:%v and inorder:%v wanted %v but got %v", tt.preorder, tt.inorder, utils.LevelOrderTravesal(tt.exp), utils.LevelOrderTravesal(out))
			}
			t.Log("pass")
		})
	}
}
