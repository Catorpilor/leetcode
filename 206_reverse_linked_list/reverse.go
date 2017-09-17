package reverse

import "github.com/catorpilor/leetcode/utils"

func Reverse(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	next := head.Next
	head.Next = nil
	for next != nil {
		temp := next.Next
		next.Next = head
		head = next
		next = temp
	}
	return head
}
