package list

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestInsertionSort(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"empty list", nil, nil},
		{"one node", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"all identical", utils.ConstructFromSlice([]int{2, 2, 2}), utils.ConstructFromSlice([]int{2, 2, 2})},
		{"testcase1", utils.ConstructFromSlice([]int{5, 4, 3, 2, 1}), utils.ConstructFromSlice([]int{1, 2, 3, 4, 5})},
		{"testcase2", utils.ConstructFromSlice([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}), utils.ConstructFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := insertionSort(tt.head)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestFasterInsertionSort(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"empty list", nil, nil},
		{"one node", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"all identical", utils.ConstructFromSlice([]int{2, 2, 2}), utils.ConstructFromSlice([]int{2, 2, 2})},
		{"testcase1", utils.ConstructFromSlice([]int{5, 4, 3, 2, 1}), utils.ConstructFromSlice([]int{1, 2, 3, 4, 5})},
		{"testcase2", utils.ConstructFromSlice([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}), utils.ConstructFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := fasterInsertSort(tt.head)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
