package mkl

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestMergeKLists(t *testing.T) {
	st := []struct {
		name  string
		lists []*utils.ListNode
		exp   *utils.ListNode
	}{
		{"empty lists", nil, nil},
		{"k = 1", []*utils.ListNode{utils.ConstructFromSlice([]int{1, 2, 3})}, utils.ConstructFromSlice([]int{1, 2, 3})},
		{"testcase1", []*utils.ListNode{utils.ConstructFromSlice([]int{1, 4, 5}), utils.ConstructFromSlice([]int{1, 3, 4}), utils.ConstructFromSlice([]int{2, 6})}, utils.ConstructFromSlice([]int{1, 1, 2, 3, 4, 4, 5, 6})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := mergeKLists(tt.lists)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with lists: %v wanted %s but got %s", tt.lists, tt.exp.String(), out.String())
			}
			t.Log("pass")
		})
	}
}
