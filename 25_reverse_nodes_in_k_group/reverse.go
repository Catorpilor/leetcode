package reverse

import (
	"github.com/catorpilor/leetcode/utils"
)

func reverseK(head *utils.ListNode, k int) *utils.ListNode {
	return iterator(head, k)
}

// iterator time complexity O(N), space complexity O(1)
func iterator(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || k <= 1 {
		return head
	}
	dummy := &utils.ListNode{Next: head}
	// count nodes
	n := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}
	groupCount := n / k
	// remainings := n % k
	cur = dummy
	for i := 0; i < groupCount; i++ {
		start, tmp := cur.Next, cur.Next
		for j := 1; j < k; j++ {
			start = start.Next
		}
		next := start.Next
		start.Next = nil
		cur.Next = utils.ReverseList(cur.Next)
		tmp.Next = next
		cur = tmp
	}
	return dummy.Next
}
