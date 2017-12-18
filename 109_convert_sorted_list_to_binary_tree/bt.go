package bt

import (
	"github.com/catorpilor/leetcode/utils"
)

func SortedListToBST(head *utils.ListNode) *utils.TreeNode {
	if head == nil {
		return nil
	}
	return helper(head, nil)

}

func helper(head, tail *utils.ListNode) *utils.TreeNode {
	slow, fast := head, head
	if head == tail {
		return nil
	}
	for fast != tail && fast.Next != tail && fast.Next.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	n := &utils.TreeNode{Val: slow.Val}
	n.Left = helper(head, slow)
	n.Right = helper(slow.Next, tail)
	return n
}

func SortedListToBST2(head *utils.ListNode) *utils.TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &utils.TreeNode{Val: head.Val}
	}
	cnt := 1
	p := head
	for p.Next != nil {
		p = p.Next
		cnt++
	}
	return inorder(head, 1, cnt)

}

func inorder(node *utils.ListNode, low, hi int) *utils.TreeNode {
	if low > hi {
		return nil
	}
	mid := low + (hi-low)/2
	left := inorder(node, low, mid-1)
	n := &utils.TreeNode{Val: node.Val}
	n.Left = left
	node = node.Next
	right := inorder(node, mid+1, hi)
	n.Right = right
	return n
}
