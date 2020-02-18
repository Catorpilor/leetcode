package pl

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestPartition(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		x    int
		exp  *utils.ListNode
	}{
		{"empty list", nil, 3, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), 2, utils.ConstructFromSlice([]int{1})},
		{"failed1", utils.ConstructFromSlice([]int{2, 1}), 2, utils.ConstructFromSlice([]int{1, 2})},
		{"all identical", utils.ConstructFromSlice([]int{2, 2, 2, 2, 2}), 3, utils.ConstructFromSlice([]int{2, 2, 2, 2, 2})},
		{"testcase1", utils.ConstructFromSlice([]int{1, 4, 3, 2, 5, 2}), 3, utils.ConstructFromSlice([]int{1, 2, 2, 4, 3, 5})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := partition(tt.head, tt.x)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s and x: %d wanted %s but got %s", tt.head, tt.x, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
