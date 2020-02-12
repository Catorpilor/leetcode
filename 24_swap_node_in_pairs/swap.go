package swap

import "github.com/catorpilor/leetcode/utils"

func swapPairs(head *utils.ListNode) *utils.ListNode {
	return swapNodes(head)
}

// swapByValue just swap values instead, which is not allowed
func swapByValue(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}
	p1, p2 := head, head.Next
	var temp int
	for p2 != nil {
		// swap adjacent
		temp = p1.Val
		p1.Val = p2.Val
		p2.Val = temp

		// update pointer
		p1 = p2.Next
		if p1 != nil && p1.Next != nil {
			p2 = p1.Next
		} else {
			p2 = nil
		}

	}

	return head
}

// swapNodes time complexity O(N), space complexity O(1)
func swapNodes(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	dummy := &utils.ListNode{Next: head}
	prev := dummy
	for head != nil && head.Next != nil {
		first, second := head, head.Next
		head = second.Next
		first.Next = head
		prev.Next = second
		second.Next = first
		prev = first
	}
	return dummy.Next
}
