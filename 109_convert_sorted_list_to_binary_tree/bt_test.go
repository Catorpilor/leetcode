package bt

import (
	"reflect"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestSortedListToBST(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		exp  *utils.TreeNode
	}{
		{"nil heads", nil, nil},
		{"single heads", &utils.ListNode{Val: 1}, &utils.TreeNode{Val: 1}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := SortedListToBST(c.head)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %v",
					c.exp, ret, c.head)
			}
		})
	}
}
