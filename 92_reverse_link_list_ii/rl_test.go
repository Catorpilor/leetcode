package rl

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestReverseList(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		m, n int
		exp  *utils.ListNode
	}{
		{"single node", utils.ConstructFromSlice([]int{1}), 1, 1, utils.ConstructFromSlice([]int{1})},
		{"reverse whole list", utils.ConstructFromSlice([]int{1, 2, 3}), 1, 3, utils.ConstructFromSlice([]int{3, 2, 1})},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3, 4, 5}), 2, 4, utils.ConstructFromSlice([]int{1, 4, 3, 2, 5})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := reverseList(tt.head, tt.m, tt.n)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with list: %s, m:%d and n:%d wanted %s but got %s", tt.head.String(), tt.m, tt.n, tt.exp.String(), out.String())
			}
			t.Log("pass")
		})
	}
}
