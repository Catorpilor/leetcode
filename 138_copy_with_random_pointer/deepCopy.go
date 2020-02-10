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

// dcopy time compleixty O(N), space complextiy O(1)
func dcopy(head *RNode) *RNode {
	if head == nil {
		return head
	}
	c := head
	// extend linked list
	// for example 1->2->3->4 after this loop becomes 1->1'->2->2'->3->3'->4->4'
	for c != nil {
		next := c.Next
		node := &RNode{Val: c.Val, Next: next}
		c.Next = node
		c = next
	}
	// update extended random links
	c = head
	for c != nil {
		if c.Random != nil {
			c.Next.Random = c.Random.Next
		}
		c = c.Next.Next
	}
	c = head
	copyHead := c.Next
	copied := copyHead
	for copied.Next != nil {
		c.Next = c.Next.Next
		c = c.Next
		copied.Next = copied.Next.Next
		copied = copied.Next
	}
	return copyHead
}
