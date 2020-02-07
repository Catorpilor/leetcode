package rl

import (
	"github.com/catorpilor/leetcode/utils"
)

func reverseList(head *utils.ListNode, m, n int) *utils.ListNode {
	return reverse(head, m, n)
}

func reverse(head *utils.ListNode, m, n int) *utils.ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}
	i := 1
	var left, prev *utils.ListNode
	dummy := &utils.ListNode{Next: head}
	prev = dummy
	left = head
	for i < m {
		prev = left
		left = left.Next
		i++
	}
	// fmt.Println(left.Val)
	prev.Next = nil
	cur := left
	next := cur.Next
	cur.Next = nil
	for next != nil && i < n {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
		i++
	}
	// fmt.Printf("before assign prev.Val: %d, head: %s\n", prev.Val, head.String())
	prev.Next = cur
	left.Next = next
	// fmt.Printf("after assign prev: %s, head: %s\n", prev.String(), head.String())
	return dummy.Next
}
