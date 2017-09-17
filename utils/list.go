package utils

import (
	"bytes"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var buf bytes.Buffer
	for l != nil {
		buf.WriteString(strconv.FormatInt(int64(l.Val), 10))
		l = l.Next
		if l != nil {
			buf.WriteString("->")
		}
	}
	return buf.String()
}

func ConstructFromSlice(s []int) *ListNode {
	var head, prev *ListNode
	for i, c := range s {
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
