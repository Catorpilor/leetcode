package list

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestDetectCycle(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		exp  *utils.ListNode
	}{
		{"no cycle", utils.ConstructFromSlice([]int{1, 2, 3}), nil},
		{"testcase1", utils.ConstructFromSliceWithCycle([]int{3, 2, 0, -4, 2}), &utils.ListNode{Val: 2}},
		{"testcase2", utils.ConstructFromSliceWithCycle([]int{1, 2, 1}), &utils.ListNode{Val: 1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := detectCycle(tt.head)
			if out.Val != tt.exp.Val {
				t.Fatalf("with input list wanted %v but got %v", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
