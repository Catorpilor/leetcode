package lc

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestHasCycle(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		exp  bool
	}{
		{"empty list", nil, false},
		{"single node", utils.ConstructFromSlice([]int{1}), false},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3, 4}), false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := hasCycle(tt.head)
			if out != tt.exp {
				t.Fatalf("with list: %s wanted %t but got %t", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
