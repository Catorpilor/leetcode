package utils

import "testing"

func TestNumsOfSccs(t *testing.T) {
	st := []struct {
		name     string
		conns    [][]int
		directed bool
		n        int
		exp      int
	}{
		{"empty", nil, false, 0, 0},
		{"single node", nil, false, 1, 1},
		{"two nodes", [][]int{[]int{0, 1}}, false, 2, 2},
		{"testcase1", [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 0}, []int{1, 3}}, false, 4, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			gp := ConstructAdjacencyList(tt.n, tt.conns, tt.directed)
			out := gp.NumofSccs()
			if out != tt.exp {
				t.Fatalf("with input conns: %v and n:%d wanted %d, but got %d", tt.conns, tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
