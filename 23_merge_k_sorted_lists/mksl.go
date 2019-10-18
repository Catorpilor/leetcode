package mksl

import (
    "container/heap"
    "math"
    "sort"

    "github.com/catorpilor/LeetCode/utils"
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
