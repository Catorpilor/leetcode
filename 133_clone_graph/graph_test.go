package graph

import "testing"

func TestCloneGraph(t *testing.T) {
	st := []struct {
		name string
		adj  [][]int
		exp  *Node
	}{
		{"empty graph", nil, nil},
		{"single node", [][]int{}, &Node{Val: 1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			ori := genNode(tt.adj)
			exp := cloneGraph(ori)
			if !isEqual(ori, exp) {
				t.Fatal("not a deep clone")
			}
			t.Log("pass")
		})
	}
}
