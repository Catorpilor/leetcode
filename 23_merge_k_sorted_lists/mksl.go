package mksl

import (
	"container/heap"
	"fmt"
	"math"
	"sort"

	"github.com/catorpilor/leetcode/utils"
)

func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
	// return bruteForce(lists)
	return compare(lists)
}

func bruteForce(lists []*utils.ListNode) *utils.ListNode {
	// Space O(N) time O(NlgN)
	nodes := make([]int, 0, 10000)
	for _, l := range lists {
		for l != nil {
			nodes = append(nodes, l.Val)
			l = l.Next
		}
	}
	if len(nodes) > 1 {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i] <= nodes[j]
		})
	}
	return utils.ConstructFromSlice(nodes)
}

func compare(lists []*utils.ListNode) *utils.ListNode {
	k := len(lists)
	pointers := make([]*utils.ListNode, k)
	for i := range pointers {
		pointers[i] = lists[i]
	}
	sortedNodes := make([]int, 0, 1000000)
	for {
		sm, idx := math.MaxInt32, -1
		for i := range pointers {
			if pointers[i] != nil && pointers[i].Val < sm {
				sm = pointers[i].Val
				idx = i
			}
		}
		if idx == -1 {
			// we are done
			break
		}
		sortedNodes = append(sortedNodes, pointers[idx].Val)
		// update pointers[idx]
		pointers[idx] = pointers[idx].Next
	}
	return utils.ConstructFromSlice(sortedNodes)
}

func priorityQueue(lists []*utils.ListNode) *utils.ListNode {
	// Time O(Nlgk) k is the number of linked lists
	// the comparision will reduced to O(lgk) for every pop and push to prioirty queue
	// get the node from pq costs O(1)
	// there are N nodes

	// Space O(N) creating a new linked lists
	pq := &utils.ListNodePriorityQueue{}
	for _, head := range lists {
		if head != nil {
			// for example [[]]
			heap.Push(pq, head)
		}
	}
	if pq.Len() < 1 {
		return nil
	}
	res := heap.Pop(pq).(*utils.ListNode)
	if res.Next != nil {
		heap.Push(pq, res.Next)
	}
	tail := res
	for pq.Len() > 0 {
		tail.Next = heap.Pop(pq).(*utils.ListNode)
		tail = tail.Next
		if tail.Next != nil {
			heap.Push(pq, tail.Next)
		}
	}
	return res
}

func divideAndConquer(lists []*utils.ListNode) *utils.ListNode {
	// Time complexity O(Nlogk) Space complexity O(1)
	n := len(lists)
	if n < 1 {
		return nil
	}
	interval := 1
	for interval < n {
		fmt.Printf("interval: %d\n", interval)
		for i := 0; i < n-interval; i += interval * 2 {
			fmt.Printf("current n: %d , i: %d, interval: %d, i+interval: %d\n", n, i, interval, i+interval)
			lists[i] = merge2Lists(lists[i], lists[i+interval])
		}
		interval *= 2
	}
	return lists[0]
}

func merge2Lists(l1, l2 *utils.ListNode) *utils.ListNode {
	head := &utils.ListNode{Val: 0}
	pointer := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pointer.Next = l1
			l1 = l1.Next
		} else {
			pointer.Next = l2
			l2 = l2.Next
			// l2 = l1
			// l1 = pointer.Next.Next
		}
		pointer = pointer.Next
	}
	if l1 != nil {
		pointer.Next = l1
	} else {
		pointer.Next = l2
	}
	return head.Next
}
