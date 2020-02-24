package split

import "github.com/catorpilor/leetcode/utils"

func splitIntoParts(root *utils.ListNode, k int) []*utils.ListNode {
	return iterator(root, k)
}

// iterator time compleixty O(N), space complexity O(K)
func iterator(root *utils.ListNode, k int) []*utils.ListNode {
	res := make([]*utils.ListNode, k)
	// traverse count the number of nodes
	n := 0
	cur := root
	for cur != nil {
		cur = cur.Next
		n++
	}
	avg := n / k
	remainings := n % k
	dummy := &utils.ListNode{Next: root}
	cur = dummy
	for i := range res {
		res[i] = cur.Next
		// we split the linked list
		count := avg
		if remainings > 0 {
			count++
			remainings--
		}
		for j := 0; j < count; j++ {
			cur = cur.Next
		}
		dummy.Next = cur.Next
		cur.Next = nil
		cur = dummy
	}
	return res
}
