package list

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestRemoveNthNode(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		n    int
		exp  *utils.ListNode
	}{
		{"empty list", nil, 0, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), 1, nil},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3, 4, 5}), 5, utils.ConstructFromSlice([]int{2, 3, 4, 5})},
		{"testcase2", utils.ConstructFromSlice([]int{3, 5, 6, 2, 4}), 2, utils.ConstructFromSlice([]int{3, 5, 6, 4})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := removeNthNode(tt.head, tt.n)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s and n: %d wanted %s but got %s", tt.head, tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
