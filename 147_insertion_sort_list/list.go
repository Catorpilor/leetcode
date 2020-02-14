package list

import "github.com/catorpilor/leetcode/utils"

func insertionSort(head *utils.ListNode) *utils.ListNode {
	return insertSort(head)
}

// insertSort time complexity O(N^2), space complexity O(1)
func insertSort(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	dummy := &utils.ListNode{}
	prev, cur := dummy, head
	for cur != nil {
		next := cur.Next
		for prev.Next != nil && prev.Next.Val < cur.Val {
			prev = prev.Next
		}
		// insert cur between prev and prev.Next
		cur.Next = prev.Next
		prev.Next = cur
		cur = next
		// update prev
		prev = dummy
	}
	return dummy.Next
}

// fasterInsertSort average time complexity O(N), worst case O(N^2)
func fasterInsertSort(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	dummy := &utils.ListNode{Next: head}
	prev, cur := head, head.Next
	for cur != nil {
		if cur.Val >= prev.Val {
			prev = cur
			cur = cur.Next
			continue
		}
		prev.Next = cur.Next
		spot := dummy
		for spot.Next.Val <= cur.Val {
			spot = spot.Next
		}
		cur.Next = spot.Next
		spot.Next = cur
		cur = prev.Next
	}
	return dummy.Next
}
