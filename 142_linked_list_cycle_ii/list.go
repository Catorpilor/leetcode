package list

import "github.com/catorpilor/leetcode/utils"

func detectCycle(head *utils.ListNode) *utils.ListNode {
	return nil
}

// useTwoPointes time complexity O(N), space complexity O(1)
func useTwoPointers(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// no cycle.
	if fast == nil || fast.Next == nil {
		return nil
	}
	// slow and fast meet somewhere (distance to the start of the cycle is `m`)in the cycle.
	// suppose fast moves k steps then fast moves 2k steps. 2k - k = nr where n is the number of cycles fast travals `r` is
	// the length of the cycle.
	// then we have s = k - m where `s` is the distance between head and start of the cycle.
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
