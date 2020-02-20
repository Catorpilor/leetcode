package ml

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestMiddleNode(t *testing.T) {
	st := []struct {
		name      string
		head, exp *utils.ListNode
	}{
		{"empty head", nil, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"odd nodes", utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{2, 3})},
		{"even nodes", utils.ConstructFromSlice([]int{1, 2, 3, 4}), utils.ConstructFromSlice([]int{3, 4})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := middleNode(tt.head)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
