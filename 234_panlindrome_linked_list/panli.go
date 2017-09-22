package panli

import (
	"github.com/catorpilor/leetcode/utils"
)

func reverseWithAlloc(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	rh := &utils.ListNode{Val: head.Val, Next: head.Next}
	next := rh.Next
	rh.Next = nil
	for next != nil {
		node := &utils.ListNode{Val: next.Val, Next: next.Next}
		temp := node.Next
		node.Next = rh
		rh = node
		next = temp
	}
	return rh
}

func reverse(head *utils.ListNode) *utils.ListNode {
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

// IsPalindromeWithAlloc is an O(n) time, O(n) space algorithm
func IsPalindromeWithAlloc(head *utils.ListNode) bool {
	// reverse linked list
	rev := reverseWithAlloc(head)
	// fmt.Printf("reversed linked list: %s\n", rev)
	// fmt.Printf("origin linked list: %s\n", head)
	for rev != nil {
		if rev.Val == head.Val {
			rev, head = rev.Next, head.Next
		} else {
			return false
		}
	}
	return true
}

func IsPalindrome(head *utils.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	slow.Next = reverse(slow.Next)
	slow = slow.Next
	for slow != nil {
		if head.Val != slow.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}
	return true
}
