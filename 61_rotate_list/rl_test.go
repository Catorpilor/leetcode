package rl

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestRotateRight(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		k    int
		exp  *utils.ListNode
	}{
		{"empty list", nil, 4, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), 4, utils.ConstructFromSlice([]int{1})},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3, 4, 5}), 2, utils.ConstructFromSlice([]int{4, 5, 1, 2, 3})},
		{"testcase2", utils.ConstructFromSlice([]int{0, 1, 2}), 4, utils.ConstructFromSlice([]int{2, 0, 1})},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := rotateRight(tt.head, tt.k)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s and k: %d wanted %s but got %s", tt.head, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestRotateRightUseIter(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		k    int
		exp  *utils.ListNode
	}{
		{"empty list", nil, 4, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), 4, utils.ConstructFromSlice([]int{1})},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3, 4, 5}), 2, utils.ConstructFromSlice([]int{4, 5, 1, 2, 3})},
		{"testcase2", utils.ConstructFromSlice([]int{0, 1, 2}), 4, utils.ConstructFromSlice([]int{2, 0, 1})},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useIter(tt.head, tt.k)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s and k: %d wanted %s but got %s", tt.head, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
