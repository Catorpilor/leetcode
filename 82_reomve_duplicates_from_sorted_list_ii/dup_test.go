package dup

import (
	"math"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestDeleteDups(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"empty list", nil, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"all identical", utils.ConstructFromSlice([]int{1, 1, 1, 1, 1}), nil},
		{"testcase1", utils.ConstructFromSlice([]int{1, 1, 2, 2, 3, 4}), utils.ConstructFromSlice([]int{3, 4})},
		{"failed1", utils.ConstructFromSlice([]int{1, 2, 3, 3, 4, 4, 5}), utils.ConstructFromSlice([]int{1, 2, 5})},
		{"failed2", utils.ConstructFromSlice([]int{math.MinInt32, 2, math.MaxInt32}), utils.ConstructFromSlice([]int{math.MinInt32, 2, math.MaxInt32})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := deleteDups(tt.head)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestDeleteDupsUseDummy(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"empty list", nil, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"all identical", utils.ConstructFromSlice([]int{1, 1, 1, 1, 1}), nil},
		{"testcase1", utils.ConstructFromSlice([]int{1, 1, 2, 2, 3, 4}), utils.ConstructFromSlice([]int{3, 4})},
		{"failed1", utils.ConstructFromSlice([]int{1, 2, 3, 3, 4, 4, 5}), utils.ConstructFromSlice([]int{1, 2, 5})},
		{"failed2", utils.ConstructFromSlice([]int{math.MinInt32, 2, math.MaxInt32}), utils.ConstructFromSlice([]int{math.MinInt32, 2, math.MaxInt32})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useDummy(tt.head)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
