package add2nums

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestAddTwoNums(t *testing.T) {
	st := []struct {
		name   string
		l1, l2 *utils.ListNode
		exp    *utils.ListNode
	}{
		{"both nil", nil, nil, nil},
		{"one is nil", nil, utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{1, 2, 3})},
		{"testcase1", utils.ConstructFromSlice([]int{2, 4, 3}), utils.ConstructFromSlice([]int{5, 6, 4}), utils.ConstructFromSlice([]int{7, 0, 8})},
		{"edgecase1", utils.ConstructFromSlice([]int{1, 9, 9, 9}), utils.ConstructFromSlice([]int{9}), utils.ConstructFromSlice([]int{0, 0, 0, 0, 1})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := addTwoNums(tt.l1, tt.l2)
			if !utils.IsEqualList(out, tt.exp) {
				t.Fatalf("with l1:%s, and l2:%s wanted %s but got %s", tt.l1.String(), tt.l2.String(), tt.exp.String(), out.String())
			}
			t.Log("pass")
		})
	}
}
