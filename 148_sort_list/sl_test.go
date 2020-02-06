package sl

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestSortList(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"nil list", nil, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"testcase1", utils.ConstructFromSlice([]int{3, 2, 1, 4}), utils.ConstructFromSlice([]int{1, 2, 3, 4})},
		{"testcase2", utils.ConstructFromSlice([]int{-1, 5, 3, 4, 0}), utils.ConstructFromSlice([]int{-1, 0, 3, 4, 5})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := sortList(tt.head)
			if !utils.IsEqualList(out, tt.exp) {
				t.Fatalf("with head: %s wanted %s but got %s", tt.head.String(), tt.exp.String(), out.String())
			}
			t.Log("pass")
		})
	}
}
