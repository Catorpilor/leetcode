package utils

import (
    "container/heap"
    "testing"
)

func TestPqLen(t *testing.T) {
    st := []struct {
        name  string
        pq    ListNodePriorityQueue
        nodes []ListNode
        exp   int
    }{
        {"empty nodes", ListNodePriorityQueue{}, nil, 0},
        {"nodes' len eq 1", ListNodePriorityQueue{}, []ListNode{ListNode{Val: 1}}, 1},
        {"nodes' len eq 2", ListNodePriorityQueue{}, []ListNode{ListNode{Val: 1}, ListNode{Val: 2}}, 2},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            for _, node := range tt.nodes {
                tt.pq.Push(&node)
            }
            out := tt.pq.Len()
            if out != tt.exp {
                t.Fatalf("priorityQueue: %v and nodes: %v wanted %d but got %d", tt.pq, tt.nodes, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestPqPop(t *testing.T) {
    st := []struct {
        name  string
        pq    ListNodePriorityQueue
        nodes []ListNode
        exp   []int
    }{
        {"nodes' len eq 1", ListNodePriorityQueue{}, []ListNode{ListNode{Val: 1}}, []int{1}},
        {"nodes' len eq 2", ListNodePriorityQueue{}, []ListNode{ListNode{Val: 1}, ListNode{Val: 2}}, []int{1, 2}},
        {"nodes' len eq 3", ListNodePriorityQueue{},
            []ListNode{ListNode{Val: 3}, ListNode{Val: 1}, ListNode{Val: 2}}, []int{1, 2, 3}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            for i := range tt.nodes {
                heap.Push(&tt.pq, &(tt.nodes[i]))
            }
            var idx int
            for tt.pq.Len() > 0 {
                item := heap.Pop(&tt.pq).(*ListNode)
                if item.Val != tt.exp[idx] {
                    t.Fatalf("priorityQueue: %v and nodes: %v wanted %d but got %d", tt.pq, tt.nodes, tt.exp[idx], item.Val)
                }
                t.Log("pass")
                idx++
            }
        })
    }
}
