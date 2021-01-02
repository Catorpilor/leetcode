package rn

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func constructWithTesting(t *testing.T, head *utils.ListNode) *Solution {
	s := Constructor(head)
	t.Cleanup(func() {
		s = nil
	})
	return s
}

func TestGetRandom(t *testing.T) {
	st := []struct {
		name  string
		arr   []int
		calls int
	}{
		{"testcase1", []int{1, 2, 3}, 5},
	}
	for _, tt := range st {
		head := utils.ConstructFromSlice(tt.arr)
		s := constructWithTesting(t, head)
		for i := 0; i < tt.calls; i++ {
			out := s.GetRandom()
			t.Log(out)
		}
	}
}
