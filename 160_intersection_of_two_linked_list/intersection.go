package intersection

type ListNode struct {
	val  int
	next *ListNode
}

func getIntersection(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	// suppose len(headA) = a+c, len(headB) = b+c

	a, b := headA, headB
	for a != b {
		// if there is an intersection
		// then a == b at a+c+b steps after a point to headB
		// and b will be b+c+a steps
		if a == nil {
			a = headB // this is awssome!
		} else {
			a = a.next
		}
		if b == nil {
			b = headA
		} else {
			b = b.next
		}
	}
	return a
}
