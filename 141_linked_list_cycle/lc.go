package lc

import "github.com/catorpilor/leetcode/utils"

func hasCycle(head *utils.ListNode) bool {
	return twoPointers(head)
}

// twoPointers time complexity O(N), space complexity O(1)
func twoPointers(head *utils.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := slow.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	return false
}
