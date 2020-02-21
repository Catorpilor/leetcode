package zerosum

import "github.com/catorpilor/leetcode/utils"

func zeroSum(head *utils.ListNode) *utils.ListNode {
	return preSum(head)
}

// preSum one pass, time complexity O(N), space complexity O(N)
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

// pprefixSum two passes, time complexity O(N), space complexity O(N)
func prefixSum(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	dummy := &utils.ListNode{Next: head}
	preMap := make(map[int]*utils.ListNode)
	preMap[0] = dummy
	pSum := 0
	// calculate the preSum, and update preMap to the latest node
	for c := dummy; c != nil; c = c.Next {
		pSum += c.Val
		preMap[pSum] = c
	}
	pSum = 0
	// iterator again, calculate the preSum and directly skip to the last occurrence of that pSum
	for c := dummy; c != nil; c = c.Next {
		pSum += c.Val
		c.Next = preMap[pSum].Next
	}
	return dummy.Next
}
