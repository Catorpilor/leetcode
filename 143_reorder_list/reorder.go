package reorder

import (
	"github.com/catorpilor/leetcode/utils"
)

func reorderList(head *utils.ListNode) *utils.ListNode {
	return twoPointers(head)
}

// twoPointers time complexity O(N), space complexity O(1)
func twoPointers(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// slow.Next point to the second half of the list
	fast = slow.Next
	// cut the loop
	slow.Next = nil
	fast = utils.ReverseList(fast)
	// fmt.Printf("head: %s, fast: %s\n", head, fast)
	dummy := &utils.ListNode{Next: head}
	// the second is less or equal to the fist half
	for fast != nil {
		tmpS, tmpH := fast.Next, head.Next
		head.Next = fast
		fast.Next = tmpH
		fast, head = tmpS, tmpH
	}
	return dummy.Next
}
