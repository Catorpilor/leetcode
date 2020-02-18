package pl

import "github.com/catorpilor/leetcode/utils"

func partition(head *utils.ListNode, x int) *utils.ListNode {
	return twoPointers(head, x)
}

func twoPointers(head *utils.ListNode, x int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &utils.ListNode{}
	smallHead := dummy
	var bigHead, bc *utils.ListNode
	sc := smallHead
	cur := head
	for cur != nil {
		tmp := cur.Next
		if cur.Val < x {
			sc.Next = cur
			sc = cur
			sc.Next = nil
		} else {
			if bigHead == nil {
				// first big one, update bigHead
				bigHead = cur
				bc = bigHead
			} else {
				bc.Next = cur
				bc = cur
			}
			bc.Next = nil
		}
		cur = tmp
	}
	sc.Next = bigHead
	return dummy.Next
}
