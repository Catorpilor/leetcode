package list

import "github.com/catorpilor/leetcode/utils"

func removeNthNode(head *utils.ListNode, n int) *utils.ListNode {
	return iterator(head, n)
}

// iterator time complexity O(N), space complexity O(1)
func iterator(head *utils.ListNode, n int) *utils.ListNode {
	if head == nil || n < 1 {
		return head
	}
	head = utils.ReverseList(head)
	dummy := &utils.ListNode{Next: head}
	prev, cur := dummy, head
	idx := 1
	for cur != nil {
		if idx == n {
			prev.Next = cur.Next
			cur.Next = nil
			break
		}
		prev = cur
		cur = cur.Next
		idx++
	}
	dummy.Next = utils.ReverseList(dummy.Next)
	return dummy.Next
}

// iteratorOneWay use two pointers, time complextity O(N), space complexity O(1)
func iteratorOneWay(head *utils.ListNode, n int) *utils.ListNode {
	if head == nil {
		return head
	}
	// two pointers
	dummy := &utils.ListNode{Next: head}
	l, r := dummy, dummy
	// l, r with n+1 gap
	for i := 1; i <= n+1; i++ {
		r = r.Next
	}
	for r != nil {
		r, l = r.Next, l.Next
	}
	// when r reaches the end, l.Next is the node to remove
	l.Next = l.Next.Next
	return dummy.Next
}
