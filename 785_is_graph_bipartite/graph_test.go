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
		{"failed1", [][]int{{}, []int{2, 4, 6}, []int{1, 4, 8, 9}, []int{7, 8}, []int{1, 2, 8, 9}, []int{6, 9}, []int{1, 5, 7, 8, 9}, []int{3, 6, 9}, []int{2, 3, 4, 6, 9}, []int{2, 4, 5, 6, 7, 8}}, false},
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
