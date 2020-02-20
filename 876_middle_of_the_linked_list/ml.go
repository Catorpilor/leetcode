package ml

import "github.com/catorpilor/leetcode/utils"

func middleNode(head *utils.ListNode) *utils.ListNode {
	return twoPointers(head)
}

func twoPointers(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil {
		return slow.Next
	}
	return slow
}
