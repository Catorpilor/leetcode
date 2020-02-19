package dup

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func deleteDups(head *utils.ListNode) *utils.ListNode {
	return useDummy(head)
}

// useDummy time complexity O(N), space complexity O(N)
func useDummy(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &utils.ListNode{Val: math.MaxInt32}
	cur := dummy
	dummy.Next = helper(dummy.Val, head)
	cur = dummy.Next
	for cur != nil {
		cur.Next = helper(cur.Val, cur.Next)
		cur = cur.Next
	}
	return dummy.Next
}

// heper return the next un-duplicated node starts from cur
func helper(prev int, cur *utils.ListNode) *utils.ListNode {
	for cur != nil && cur.Val == prev {
		cur = cur.Next
	}
	if cur == nil || cur.Next == nil || cur.Val != cur.Next.Val {
		return cur
	}
	return helper(cur.Val, cur.Next)
}
