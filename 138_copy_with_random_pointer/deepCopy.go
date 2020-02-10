package deepCopy

import "math"

// RNode is a ListNode with a random pointer
type RNode struct {
	Val    int
	Next   *RNode
	Random *RNode
}

func constructFromSlice(list [][2]int) *RNode {
	n := len(list)
	nodes := make([]*RNode, 0, n)
	for i := range list {
		node := &RNode{Val: list[i][0]}
		nodes = append(nodes, node)
	}
	dummy := &RNode{}
	cur := dummy
	for i := range list {
		cur.Next = nodes[i]
		if list[i][1] != math.MinInt32 {
			cur.Random = nodes[list[i][1]]
		}
		cur = cur.Next
	}
	return dummy.Next
}

func isEqualRNode(l1, l2 *RNode) bool {
	if (l1 == nil && l2 != nil) || (l1 != nil && l2 == nil) {
		return false
	}
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil || l2 != nil {
		return false
	}

	return true
}

func dcopy(head *RNode) *RNode {
	if head == nil {
		return head
	}
	return nil
}
