package utils

import (
	"bytes"
	"strconv"
)

// ListNode is a node structure in a list
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

// IsEqualList checks l1 == l2
func IsEqualList(h1, h2 *ListNode) bool {
	return h1.String() == h2.String()
}

// ConstructFromSlice returns a *ListNode built by s
func ConstructFromSlice(s []int) *ListNode {
	// var head, prev *ListNode
	// for i, c := range s {
	// 	if i == 0 {
	// 		head = &ListNode{
	// 			Val:  c,
	// 			Next: nil,
	// 		}
	// 		prev = head
	// 	} else {
	// 		cur := &ListNode{
	// 			Val:  c,
	// 			Next: nil,
	// 		}
	// 		if prev != nil {
	// 			prev.Next = cur
	// 		}
	// 		prev = cur
	// 	}
	// }
	// return head
	dummy := &ListNode{}
	pre := dummy
	for i := range s {
		pre.Next = &ListNode{Val: s[i]}
		pre = pre.Next
	}
	return dummy.Next
}

// ReverseList return the reversed  list
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	next := head.Next
	head.Next = nil
	for next != nil {
		tmp := next.Next
		next.Next = head
		head = next
		next = tmp
	}
	return head
}
