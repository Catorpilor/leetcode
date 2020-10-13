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

// useIter time complexity O(N), space complexity O(1)
func useIter(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oldTail := head
	n := 1
	for oldTail.Next != nil {
		oldTail = oldTail.Next
		n++
	}
	// make it a ring
	oldTail.Next = head
	// new tail at pos n - k%n -1
	newTail := head
	for i := 0; i < n-k%n-1; i++ {
		newTail = newTail.Next
	}
	// newHead at pos n - k%n
	newHead := newTail.Next
	// break the ring
	newTail.Next = nil
	return newHead
}

// useReverseV2 time complexity O(N), space complexity O(1)
func useReverseV2(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var n int
	tmp := head
	for tmp != nil {
		n++
		tmp = tmp.Next
	}
	k %= n
	if k == 0 {
		return head
	}
	// first reverse
	tail := utils.ReverseList(head)
	var prev *utils.ListNode
	oldT := tail
	for i := 1; i <= k; i++ {
		tmp := oldT.Next
		oldT.Next = prev
		prev = oldT
		oldT = tmp
	}
	tail.Next = utils.ReverseList(oldT)
	return prev
}
