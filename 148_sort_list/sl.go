package sl

import "github.com/catorpilor/leetcode/utils"

func sortList(head *utils.ListNode) *utils.ListNode {
	return bottomUp(head)
}

// bottomUp time complexity O(nlgn), space complexity O(1)
func bottomUp(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// count list length
	var n int
	tmp := head
	for tmp != nil {
		tmp = tmp.Next
		n++
	}
	// for example with list head: 3->2->1->4
	dummy := &utils.ListNode{Next: head}
	// becomes dummy->3->2->1->4
	var left, right, tail *utils.ListNode
	for i := 1; i < n; i <<= 1 {
		cur := dummy.Next
		tail = dummy
		for cur != nil {
			left = cur
			// left: 3
			right = split(left, i)
			// right: 2
			cur = split(right, i)
			// cur: 1
			tail = merge(left, right, tail)
			// tail is 3
		}
	}
	return dummy.Next
}

// split split the list into two lists, h contains n nodes, returns the second list
func split(h *utils.ListNode, n int) *utils.ListNode {
	for i := 1; h != nil && i < n; i++ {
		h = h.Next
	}
	if h == nil {
		return h
	}
	nxt := h.Next
	h.Next = nil
	return nxt
}

// merge merge two sorted listnode to head
func merge(l1, l2, head *utils.ListNode) *utils.ListNode {
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = l2
			cur = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			cur = l1
			l1 = l1.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}
