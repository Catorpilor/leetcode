package sort

import (
	"reflect"
	"testing"
)

type res struct {
	l1  *ListNode
	l2  *ListNode
	exp *ListNode
}

func buildWithSlice(a []int) *ListNode {
	var h, c *ListNode
	for _, v := range a {
		n := &ListNode{v, nil}
		if c != nil {
			c.Next = n
		}
		if h == nil {
			h = n
		}
		c = n
	}
	return h
}

func TestMergeTwoLists(t *testing.T) {
	cases := []res{
		res{nil, nil, nil},
		res{nil, buildWithSlice([]int{5}), buildWithSlice([]int{5})},
		res{buildWithSlice([]int{1, 4, 6}), buildWithSlice([]int{5}), buildWithSlice([]int{1, 4, 5, 6})},
	}

	for _, c := range cases {
		ret := MergeTwoLists(c.l1, c.l2)
		if !reflect.DeepEqual(ret, c.exp) {
			t.Errorf("expected %v but got %v with input %v,%v", c.exp, ret, c.l1, c.l2)
		}
	}
	t.Log("passed")
}
