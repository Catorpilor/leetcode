package add

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestAddTwo(t *testing.T) {
	st := []struct {
		name        string
		l1, l2, exp *utils.ListNode
	}{
		{"both empty", nil, nil, nil},
		{"l1 is empty", nil, utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{1, 2, 3})},
		{"testcase1", utils.ConstructFromSlice([]int{7, 2, 4, 3}), utils.ConstructFromSlice([]int{5, 6, 4}), utils.ConstructFromSlice([]int{7, 8, 0, 7})},
		{"testcase2", utils.ConstructFromSlice([]int{9, 9, 9, 9, 9}), utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1, 0, 0, 0, 0, 0})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := addTwo(tt.l1, tt.l2)
			if !utils.IsEqualList(out, tt.exp) {
				t.Fatalf("with input l1: %s and l2: %s wanted %s but got %s", tt.l1, tt.l2, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestReverse(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"single node", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{3, 2, 1})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := reverse(tt.head)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with head: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestUseStack(t *testing.T) {
	st := []struct {
		name        string
		l1, l2, exp *utils.ListNode
	}{
		{"both empty", nil, nil, nil},
		{"l1 is empty", nil, utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{1, 2, 3})},
		{"testcase1", utils.ConstructFromSlice([]int{7, 2, 4, 3}), utils.ConstructFromSlice([]int{5, 6, 4}), utils.ConstructFromSlice([]int{7, 8, 0, 7})},
		{"testcase3", utils.ConstructFromSlice([]int{0, 2, 4, 3}), utils.ConstructFromSlice([]int{5, 6, 4}), utils.ConstructFromSlice([]int{8, 0, 7})},
		{"testcase2", utils.ConstructFromSlice([]int{9, 9, 9, 9, 9}), utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1, 0, 0, 0, 0, 0})},
		{"failed1", utils.ConstructFromSlice([]int{0}), utils.ConstructFromSlice([]int{0}), utils.ConstructFromSlice([]int{0})},
		{"failed2", utils.ConstructFromSlice([]int{0, 0, 1}), utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{2})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useStack(tt.l1, tt.l2)
			if !utils.IsEqualList(out, tt.exp) {
				t.Fatalf("with input l1: %s and l2: %s wanted %s but got %s", tt.l1, tt.l2, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
