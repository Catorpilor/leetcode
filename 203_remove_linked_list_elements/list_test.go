package list

import (
	"reflect"
	"testing"
)

func constructListFromSlice(vals []int) *ListNode {
	var head, prev *ListNode
	for i, c := range vals {
		if i == 0 {
			head = &ListNode{
				Val:  c,
				Next: nil,
			}
			prev = head
		} else {
			cur := &ListNode{
				Val:  c,
				Next: nil,
			}
			if prev != nil {
				prev.Next = cur
			}
			prev = cur
		}
	}
	return head
}

func TestRemoveElements(t *testing.T) {
	st := []struct {
		name string
		h    *ListNode
		val  int
		exp  *ListNode
	}{
		{"nil list", nil, 1, nil},
		{"single node with Val eq val", constructListFromSlice([]int{1}), 1, nil},
		{"single node with Val neq val", constructListFromSlice([]int{2}), 1, constructListFromSlice([]int{2})},
		{"normal case", constructListFromSlice([]int{1, 2, 6, 3, 4, 5, 6}), 6, constructListFromSlice([]int{1, 2, 3, 4, 5})},
		{"extram case", constructListFromSlice([]int{2, 2, 2, 2, 2, 2, 2}), 2, nil},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := RemoveElements2(c.h, c.val)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %s but got %s, with input %s and %d",
					c.exp, ret, c.h, c.val)
			}
		})
	}
}
