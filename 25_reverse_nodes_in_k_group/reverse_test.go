package reverse

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestReverseKGroup(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		k    int
		exp  *utils.ListNode
	}{
		{"head is nill", nil, 1, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), 1, utils.ConstructFromSlice([]int{1})},
		{"reverse n", utils.ConstructFromSlice([]int{1, 2, 3, 4}), 4, utils.ConstructFromSlice([]int{4, 3, 2, 1})},
		{"testcase1", utils.ConstructFromSlice([]int{1, 2, 3, 4, 5}), 2, utils.ConstructFromSlice([]int{2, 1, 3, 4, 5})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := reverseK(tt.head, tt.k)
			if !utils.IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
