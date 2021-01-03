package tree

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestGetTargetCopy(t *testing.T) {
	st := []struct {
		name                    string
		orignal, cloned, target *utils.TreeNode
		exp                     *utils.TreeNode
	}{
		{"testcaes1", utils.ConstructTree([]int{7, 4, 3, math.MinInt32, math.MinInt32, 6, 19}),
			utils.ConstructTree([]int{7, 4, 3, math.MinInt32, math.MinInt32, 6, 19}), &utils.TreeNode{Val: 3}, &utils.TreeNode{Val: 3}},
		{"testcase2", utils.ConstructTree([]int{7}), utils.ConstructTree([]int{7}), &utils.TreeNode{Val: 7}, &utils.TreeNode{Val: 7}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := getTargetCopy(tt.orignal, tt.cloned, tt.target)
			if out == nil || out.Val != tt.exp.Val {
				t.Fatalf("wanted %v but got %v", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
