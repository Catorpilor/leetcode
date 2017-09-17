package list

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

func RemoveElements(head *ListNode, val int) *ListNode {
	prev, cur := head, head
	for cur != nil {
		if cur.Val == val {
			if cur == head {
				head = head.Next
				cur = cur.Next
				prev = cur
			} else {
				prev.Next = cur.Next
				cur = cur.Next
			}
		} else {
			prev = cur
			cur = cur.Next
		}

	}
	return head
}

func RemoveElements2(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return dummy.Next
}
