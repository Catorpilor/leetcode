package utils

import (
	"container/heap"
	"testing"
)

func TestTuplePqLen(t *testing.T) {
	st := []struct {
		name  string
		pq    TuplePriorityQueue
		nodes [][3]int
		exp   int
	}{
		{"empty nodes", TuplePriorityQueue{}, nil, 0},
		{"nodes' len eq 1", TuplePriorityQueue{}, [][3]int{[3]int{1, 2, 3}, [3]int{2, 2, 4}}, 2},
		{"nodes' len eq 2", TuplePriorityQueue{}, [][3]int{[3]int{2, 2, 2}}, 1},
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

func TestTuplePqPop(t *testing.T) {
	st := []struct {
		name  string
		pq    TuplePriorityQueue
		nodes [][3]int
		exp   [][3]int
	}{
		{"nodes' len eq 1", TuplePriorityQueue{}, [][3]int{[3]int{1, 2, 3}}, [][3]int{[3]int{1, 2, 3}}},
		{"nodes' len eq 2", TuplePriorityQueue{}, [][3]int{[3]int{1, 2, 3}, [3]int{2, 2, 2}}, [][3]int{[3]int{1, 2, 3}, [3]int{2, 2, 2}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			for i := range tt.nodes {
				heap.Push(&tt.pq, tt.nodes[i])
			}
			var idx int
			for tt.pq.Len() > 0 {
				item := heap.Pop(&tt.pq).([3]int)
				if item != tt.exp[idx] {
					t.Fatalf("priorityQueue: %v and nodes: %v wanted %d but got %d", tt.pq, tt.nodes, tt.exp[idx], item)
				}
				t.Log("pass")
				idx++
			}
		})
	}
}

func TestTuplePqTop(t *testing.T) {
	st := []struct {
		name  string
		pq    TuplePriorityQueue
		nodes [][3]int
		exp   [3]int
	}{
		{"nodes' len eq 1", TuplePriorityQueue{}, [][3]int{[3]int{1, 2, 3}}, [3]int{1, 2, 3}},
		{"nodes' len eq 2", TuplePriorityQueue{}, [][3]int{[3]int{1, 2, 3}, [3]int{2, 3, 4}}, [3]int{1, 2, 3}},
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
