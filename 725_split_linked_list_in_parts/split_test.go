package split

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestSplitIntoParts(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		k    int
		exp  []*utils.ListNode
	}{
		{"empty head", nil, 3, []*utils.ListNode{nil, nil, nil}},
		{"single node", utils.ConstructFromSlice([]int{1}), 3, []*utils.ListNode{utils.ConstructFromSlice([]int{1}), nil, nil}},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3, 4, 5}), 2, []*utils.ListNode{utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{4, 5})}},
		{"testcase2", utils.ConstructFromSlice([]int{1, 2, 3, 4, 5}), 1, []*utils.ListNode{utils.ConstructFromSlice([]int{1, 2, 3, 4, 5})}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			outs := splitIntoParts(tt.head, tt.k)
			if len(outs) != len(tt.exp) {
				t.Fatalf("len(outs) != len(tt.exp), wanted %d but got %d", len(tt.exp), len(outs))
			}
			for i := range tt.exp {
				if !utils.IsEqualList(tt.exp[i], outs[i]) {
					t.Fatalf("wanted %s but got %s", tt.exp[i], outs[i])
				}
			}
			t.Log("pass")
		})
	}
}
