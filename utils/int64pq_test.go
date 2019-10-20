package utils

import (
    "container/heap"
    "testing"
)

func TestInt64PqLen(t *testing.T) {
    st := []struct {
        name  string
        pq    Int64PriorityQueue
        nodes []int64
        exp   int
    }{
        {"empty nodes", Int64PriorityQueue{}, nil, 0},
        {"nodes' len eq 1", Int64PriorityQueue{}, []int64{1}, 1},
        {"nodes' len eq 2", Int64PriorityQueue{}, []int64{1, 2}, 2},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            for _, node := range tt.nodes {
                tt.pq.Push(node)
            }
            out := tt.pq.Len()
            if out != tt.exp {
                t.Fatalf("priorityQueue: %v and nodes: %v wanted %d but got %d", tt.pq, tt.nodes, tt.exp, out)
            }
            t.Log("pass")
        })
    }
}

func TestInt64PqPop(t *testing.T) {
    st := []struct {
        name  string
        pq    Int64PriorityQueue
        nodes []int64
        exp   []int64
    }{
        {"nodes' len eq 1", Int64PriorityQueue{}, []int64{1}, []int64{1}},
        {"nodes' len eq 2", Int64PriorityQueue{}, []int64{1, 2}, []int64{1, 2}},
        {"nodes' len eq 3", Int64PriorityQueue{},
            []int64{3, 1, 2}, []int64{1, 2, 3}},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            for i := range tt.nodes {
                heap.Push(&tt.pq, tt.nodes[i])
            }
            var idx int
            for tt.pq.Len() > 0 {
                item := heap.Pop(&tt.pq).(int64)
                if item != tt.exp[idx] {
                    t.Fatalf("priorityQueue: %v and nodes: %v wanted %d but got %d", tt.pq, tt.nodes, tt.exp[idx], item)
                }
                t.Log("pass")
                idx++
            }
        })
    }
}

func TestInt64PqTop(t *testing.T) {
    st := []struct {
        name  string
        pq    Int64PriorityQueue
        nodes []int64
        exp   int64
    }{
        {"nodes' len eq 1", Int64PriorityQueue{}, []int64{1}, 1},
        {"nodes' len eq 2", Int64PriorityQueue{}, []int64{1, 2}, 1},
        {"nodes' len eq 3", Int64PriorityQueue{},
            []int64{3, 1, 2}, 1},
    }
    for _, tt := range st {
        t.Run(tt.name, func(t *testing.T) {
            for i := range tt.nodes {
                heap.Push(&tt.pq, tt.nodes[i])
            }
            item := tt.pq.Top()
            if item != tt.exp {
                t.Fatalf("priorityQueue: %v and nodes: %v wanted %d but got %d", tt.pq, tt.nodes, tt.exp, item)
            }
            t.Log("pass")
        })
    }
}
