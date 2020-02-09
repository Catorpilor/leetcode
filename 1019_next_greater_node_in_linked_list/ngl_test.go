package ngl

import (
	"reflect"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestNextLargeNode(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		exp  []int
	}{
		{"nil head", nil, []int{}},
		{"single node", utils.ConstructFromSlice([]int{1}), []int{0}},
		{"identical", utils.ConstructFromSlice([]int{1, 1, 1}), []int{0, 0, 0}},
		{"testcase1", utils.ConstructFromSlice([]int{2, 1, 5}), []int{5, 5, 0}},
		{"testcase2", utils.ConstructFromSlice([]int{2, 7, 4, 3, 5}), []int{7, 0, 5, 5, 0}},
		{"testcase3", utils.ConstructFromSlice([]int{1, 7, 5, 1, 9, 2, 5, 1}), []int{7, 9, 9, 9, 0, 5, 0, 0}},
		{"failed1", utils.ConstructFromSlice([]int{4, 3, 2, 5, 1, 8, 10}), []int{5, 5, 5, 8, 8, 10, 0}},
		{"failed2", utils.ConstructFromSlice([]int{9, 7, 6, 7, 6, 9}), []int{0, 9, 7, 9, 9, 0}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := nextLargeNode(tt.head)
			if !reflect.DeepEqual(out, tt.exp) {
				t.Fatalf("with input list: %s, wanted %v but got %v", tt.head.String(), tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
