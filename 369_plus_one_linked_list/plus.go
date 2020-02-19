package plus

import "github.com/catorpilor/leetcode/utils"

func plusOne(head *utils.ListNode) *utils.ListNode {
	return useReverse(head)
}

func useReverse(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	head = utils.ReverseList(head)
	dummy := &utils.ListNode{Next: head}
	cur := dummy
	flag := false
	for cur.Next != nil {
		cur = cur.Next
		if cur.Val < 9 {
			cur.Val++
			flag = false
			break
		}
		cur.Val = 0
		flag = true
	}
	if flag {
		cur.Next = &utils.ListNode{Val: 1}
	}
	dummy.Next = utils.ReverseList(dummy.Next)
	return dummy.Next
}
