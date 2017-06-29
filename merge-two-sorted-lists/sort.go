package sort

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret, last *ListNode
	min := make(chan *ListNode)
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
