package sort

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestMergeTwoLists(t *testing.T) {
	st := []struct {
		name        string
		l1, l2, exp *utils.ListNode
	}{
		{"both empty", nil, nil, nil},
		{"one is nil", nil, utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{1, 2, 3})},
		{"testcase1", utils.ConstructFromSlice([]int{1, 4, 6}), utils.ConstructFromSlice([]int{5}), utils.ConstructFromSlice([]int{1, 4, 5, 6})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := MergeTwoLists(tt.l1, tt.l2)
			if !utils.IsEqualList(out, tt.exp) {
				t.Fatalf("with l1: %s, and l2: %s wanted %s but got %s", tt.l1.String(), tt.l2.String(), tt.exp.String(), out.String())
			}
			t.Log("pass")
		})
	}
}
