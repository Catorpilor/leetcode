package graph

import "testing"

func TestIsBipartite(t *testing.T) {
	st := []struct {
		name  string
		graph [][]int
		exp   bool
	}{
		{"testcase1", [][]int{[]int{1, 2, 3}, []int{0, 2}, []int{0, 1, 3}, []int{0, 2}}, false},
		{"testcase2", [][]int{[]int{1, 3}, []int{0, 2}, []int{1, 3}, []int{0, 2}}, true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isBipartite(tt.graph)
			if out != tt.exp {
				t.Fatalf("wanted %t but got %t", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
