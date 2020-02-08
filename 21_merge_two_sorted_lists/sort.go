package sort

import "github.com/catorpilor/leetcode/utils"

func MergeTwoLists(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	var ret, last *utils.ListNode
	min := make(chan *utils.ListNode)
	go func() {
		defer close(min)
		for {
			if l1 == nil {
				min <- l2
				return
			}
			if l2 == nil {
				min <- l1
				return
			}
			if l1.Val < l2.Val {
				min <- l1
				l1 = l1.Next
			} else {
				min <- l2
				l2 = l2.Next
			}
		}
	}()
	for v := range min {
		if last != nil {
			last.Next = v
		}
		if ret == nil {
			// first node
			ret = v
		}
		last = v
	}
	return ret

}
