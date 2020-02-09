package mkl

import "github.com/catorpilor/leetcode/utils"

func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
	return bruteForce(lists)
}

func bruteForce(lists []*utils.ListNode) *utils.ListNode {
	n := len(lists)
	if n < 1 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	tmp := make([]*utils.ListNode, 0, n)
	for i := 0; i < n; i += 2 {
		if i+1 < n {
			tmp = append(tmp, merge(lists[i], lists[i+1]))
		} else {
			tmp = append(tmp, lists[i])
		}
	}
	for len(tmp) > 1 {
		local := make([]*utils.ListNode, 0, n)
		for i := 0; i < len(tmp); i += 2 {
			if i+1 < len(tmp) {
				local = append(local, merge(tmp[i], tmp[i+1]))
			} else {
				local = append(local, tmp[i])
			}
		}
		tmp = local
	}
	return tmp[0]
}

func merge(l1, l2 *utils.ListNode) *utils.ListNode {
	dummy := &utils.ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}
