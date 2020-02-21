package zerosum

import "github.com/catorpilor/leetcode/utils"

func zeroSum(head *utils.ListNode) *utils.ListNode {
	return preSum(head)
}

// preSum time complexity O(N), space complexity O(N)
func preSum(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	dummy := &utils.ListNode{Next: head}
	cur := dummy
	pmap := make(map[int]*utils.ListNode)
	prefix := 0
	for cur != nil {
		prefix += cur.Val
		if node, exists := pmap[prefix]; exists {
			beg := node.Next
			p := prefix + beg.Val
			for p != prefix {
				// cause we are going to delete this node, pmap has to be updated
				delete(pmap, p)
				beg = beg.Next
				p += beg.Val
			}
			node.Next = cur.Next
		} else {
			pmap[prefix] = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}
