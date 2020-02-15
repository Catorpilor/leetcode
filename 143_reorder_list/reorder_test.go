package reorder

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestReorderList(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"empty list", nil, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"two nodes", utils.ConstructFromSlice([]int{1, 3}), utils.ConstructFromSlice([]int{1, 3})},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3, 4}), utils.ConstructFromSlice([]int{1, 4, 2, 3})},
		{"testcase2", utils.ConstructFromSlice([]int{1, 2, 3, 4, 5}), utils.ConstructFromSlice([]int{1, 5, 2, 4, 3})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := reorderList(tt.head)
			// var out *utils.ListNode
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
