package mht

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindMinHeightTrees(t *testing.T) {
	st := []struct {
		name  string
		n     int
		edges [][]int
		exp   []int
	}{
		{"n=1", 1, nil, []int{0}},
		{"n=2", 2, [][]int{[]int{1, 0}}, []int{0, 1}},
		{"testcase1", 6, [][]int{[]int{3, 0}, []int{3, 1}, []int{3, 2}, []int{3, 4}, []int{4, 5}}, []int{3, 4}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findMinHeightTrees(tt.n, tt.edges)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
