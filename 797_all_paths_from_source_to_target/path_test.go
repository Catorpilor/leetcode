package path

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAllPaths(t *testing.T) {
	st := []struct {
		name  string
		graph [][]int
		exp   [][]int
	}{
		{"testcase1", [][]int{[]int{1, 2}, []int{3}, []int{3}, []int{}}, [][]int{[]int{0, 1, 3}, []int{0, 2, 3}}},
		{"testcase2", [][]int{[]int{1}, []int{2}, []int{3}, []int{}}, [][]int{[]int{0, 1, 2, 3}}},
		{"failed1", [][]int{[]int{4, 3, 1}, []int{3, 2, 4}, []int{3}, []int{4}, []int{}}, [][]int{[]int{0, 4}, []int{0, 3, 4}, []int{0, 1, 3, 4}, []int{0, 1, 2, 3, 4}, [e]int{0, 1, 4}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := allPaths(tt.graph)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("with input (-wanted, +got), %s", diff)
			}
			t.Log("pass")
		})
	}
}
