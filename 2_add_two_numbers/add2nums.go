package add2nums

import "github.com/catorpilor/leetcode/utils"

func addTwoNums(l1, l2 *utils.ListNode) *utils.ListNode {
	return iter(l1, l2)
}

func iter(l1, l2 *utils.ListNode) *utils.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	h1, h2 := l1, l2
	prev1 := h1
	carrier := false
	for h1 != nil && h2 != nil {
		h1.Val += h2.Val
		if carrier {
			h1.Val++
			carrier = false
		}
		if h1.Val >= 10 {
			h1.Val %= 10
			carrier = true
		}
		prev1 = h1
		h1, h2 = h1.Next, h2.Next
	}
	if h1 == nil {
		prev1.Next = h2
		h1 = h2
	}
	for h1 != nil {
		if !carrier {
			break
		}
		h1.Val++
		carrier = false
		if h1.Val >= 10 {
			h1.Val %= 10
			carrier = true
		}
		prev1 = h1
		h1 = h1.Next
	}
	if carrier {
		prev1.Next = &utils.ListNode{Val: 1}
	}
	return l1
}
