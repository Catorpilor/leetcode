package list

import (
	"github.com/catorpilor/leetcode/utils"
)

func oddEvenList(head *utils.ListNode) *utils.ListNode {
	return iterator(head)
}

// iterator iterates the list, time complexity O(N), space complexity O(1)
func iterator(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	dummy := &utils.ListNode{Next: head}
	oddHead, evenHead := head, head.Next
	cur := oddHead
	// head point to evenHead
	head = head.Next
	for cur != nil && head != nil {
		cur.Next = head.Next
		if cur.Next == nil {
			break
		}
		cur = cur.Next
		head.Next = cur.Next
		head = head.Next
	}
	cur.Next = evenHead
	return dummy.Next
}
