package zerosum

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestZeroSum(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"empty list", nil, nil},
		{"single node not eq 0", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"single node eq 0", utils.ConstructFromSlice([]int{0}), nil},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, -3, 3, 1}), utils.ConstructFromSlice([]int{3, 1})},
		{"testcase2", utils.ConstructFromSlice([]int{1, 2, 3, -3, 4}), utils.ConstructFromSlice([]int{1, 2, 4})},
		{"testcase3", utils.ConstructFromSlice([]int{1, 2, 3, -3, -2}), utils.ConstructFromSlice([]int{1})},
		{"testcase4", utils.ConstructFromSlice([]int{1, 3, 2, -3, -2, 5, 100, -100, 1}), utils.ConstructFromSlice([]int{1, 5, 1})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := zeroSum(tt.head)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
func TestZeroSumTwoPass(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"empty list", nil, nil},
		{"single node not eq 0", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"single node eq 0", utils.ConstructFromSlice([]int{0}), nil},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, -3, 3, 1}), utils.ConstructFromSlice([]int{3, 1})},
		{"testcase2", utils.ConstructFromSlice([]int{1, 2, 3, -3, 4}), utils.ConstructFromSlice([]int{1, 2, 4})},
		{"testcase3", utils.ConstructFromSlice([]int{1, 2, 3, -3, -2}), utils.ConstructFromSlice([]int{1})},
		{"testcase4", utils.ConstructFromSlice([]int{1, 3, 2, -3, -2, 5, 100, -100, 1}), utils.ConstructFromSlice([]int{1, 5, 1})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := prefixSum(tt.head)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
