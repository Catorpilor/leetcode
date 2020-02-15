package add

import "github.com/catorpilor/leetcode/utils"

func addTwo(l1, l2 *utils.ListNode) *utils.ListNode {
	return useReverse(l1, l2)
}

// useReverse time complexity O(M+N), space complexity O(N)
func useReverse(l1, l2 *utils.ListNode) *utils.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return nil
	}

	// reverse l1 & l2
	l1, l2 = reverse(l1), reverse(l2)
	dummy := &utils.ListNode{}
	cur := dummy
	carrier := 0
	for l1 != nil || l2 != nil || carrier != 0 {
		sum := carrier
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carrier = 0
		if sum >= 10 {
			sum %= 10
			carrier = 1
		}
		node := &utils.ListNode{Val: sum}
		cur.Next = node
		cur = node
	}
	return reverse(dummy.Next)
}

func reverse(head *utils.ListNode) *utils.ListNode {
	dummy := &utils.ListNode{}
	cur := head
	var pre, next *utils.ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	dummy.Next = pre
	return dummy.Next
}
