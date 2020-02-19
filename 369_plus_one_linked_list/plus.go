package plus

import "github.com/catorpilor/leetcode/utils"

func plusOne(head *utils.ListNode) *utils.ListNode {
	return useReverse(head)
}

// useReverse time complexity O(N), space complexity O(1)
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

// useDummy time complexity O(N), space complexit O(1)
func useDummy(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	dummy := &utils.ListNode{Next: head}
	notNine := dummy
	// find the rightmost not 9 node
	for head != nil {
		if head.Val != 9 {
			notNine = head
		}
		head = head.Next
	}
	notNine.Val++
	notNine = notNine.Next
	// update the following 9 to 0
	for notNine != nil {
		notNine.Val = 0
		notNine = notNine.Next
	}
	if dummy.Val != 0 {
		return dummy
	}
	return dummy.Next
}
