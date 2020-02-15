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

// useStack time complexity O(M+N), space complexity O(M+N)
func useStack(l1, l2 *utils.ListNode) *utils.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	a1, a2 := make([]int, 0, 65536), make([]int, 0, 65536)
	for l1 != nil {
		a1 = append(a1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		a2 = append(a2, l2.Val)
		l2 = l2.Next
	}
	n1, n2 := len(a1), len(a2)
	sum := 0
	dummy := &utils.ListNode{}
	pre := dummy.Next
	for n1 != 0 || n2 != 0 || sum != 0 {
		if n1 > 0 {
			sum += a1[n1-1]
			a1 = a1[:n1-1]
		}
		if n2 > 0 {
			sum += a2[n2-1]
			a2 = a2[:n2-1]
		}
		node := &utils.ListNode{Val: sum % 10, Next: pre}
		pre = node
		n1, n2 = len(a1), len(a2)
		sum /= 10
	}
	// skip 0 at the highest bit
	// if pre.Val == 0 {
	// 	pre = pre.Next
	// }
	dummy.Next = pre
	return dummy.Next

}
