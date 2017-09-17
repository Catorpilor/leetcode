package reverse

import (
	"reflect"
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

// func constructFromSlice(s []int) *ListNode {
// 	var head, prev *ListNode
// 	for i, c := range s {
// 		if i == 0 {
// 			head = &ListNode{
// 				Val:  c,
// 				Next: nil,
// 			}
// 			prev = head
// 		} else {
// 			cur := &ListNode{
// 				Val:  c,
// 				Next: nil,
// 			}
// 			if prev != nil {
// 				prev.Next = cur
// 			}
// 			prev = cur
// 		}
// 	}
// 	return head
// }

func TestReverse(t *testing.T) {
	st := []struct {
		name string
		head *utils.ListNode
		exp  *utils.ListNode
	}{
		{"nil list", nil, nil},
		{"single element", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"multiple element", utils.ConstructFromSlice([]int{1, 2, 3}), utils.ConstructFromSlice([]int{3, 2, 1})},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Reverse(c.head)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expeced %s but got %s, with inputs %s",
					c.exp, ret, c.head)
			}
		})
	}
}
