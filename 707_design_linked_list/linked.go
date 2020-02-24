package linked

import (
	"github.com/catorpilor/leetcode/utils"
)

// MyLinkedList wraps utils.ListNode
type MyLinkedList struct {
	Head *utils.ListNode
	N    int
}

func (ll *MyLinkedList) String() string {
	if ll.Head.Next != nil {
		return ll.Head.Next.String()
	}
	return ""
}

// Constructor creates MyLinkedList
func Constructor() *MyLinkedList {
	return &MyLinkedList{Head: &utils.ListNode{}}
}

// Get returns the ith node'value
func (ll *MyLinkedList) Get(index int) int {
	if index < 0 || index >= ll.N {
		return -1
	}
	var idx int
	cur := ll.Head
	for idx <= index {
		cur = cur.Next
		idx++
	}
	return cur.Val
}

// AddAtHead add a new head
func (ll *MyLinkedList) AddAtHead(val int) {
	ll.AddAtIndex(0, val)
}

// AddAtTail add a new node at the end
func (ll *MyLinkedList) AddAtTail(val int) {
	ll.AddAtIndex(ll.N, val)
}

// AddAtIndex at a new node at index with val, time complelxity O(N)
func (ll *MyLinkedList) AddAtIndex(index int, val int) {
	if index > ll.N || index < 0 {
		return
	}
	ll.N++
	prev := ll.Head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	node := &utils.ListNode{Val: val, Next: prev.Next}
	prev.Next = node
}

// DeleteAtIndex delete a node at index
func (ll *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= ll.N {
		return
	}
	ll.N--
	prev := ll.Head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
}
