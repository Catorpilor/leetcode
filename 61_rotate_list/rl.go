package rl

import (
	"github.com/catorpilor/leetcode/utils"
)

func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	return useReverse(head, k)
}

// useReverse time complexity O(N), space complexity O(1)
func useReverse(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}
	k %= n
	dummy := &utils.ListNode{Next: head}
	tail := utils.ReverseList(head)
	// fmt.Printf("head: %s, tail:%s, k: %d, n: %d\n", head, tail, k, n)
	for i := 1; i <= k; i++ {
		tmp := tail.Next
		head.Next = tail
		head = head.Next
		head.Next = nil
		tail = tmp
	}
	dummy.Next = utils.ReverseList(tail)
	return dummy.Next
}
