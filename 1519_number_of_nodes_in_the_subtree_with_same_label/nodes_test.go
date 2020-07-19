package nodes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumberOfNodes(t *testing.T) {
	st := []struct {
		name   string
		n      int
		edges  [][]int
		labels string
		exp    []int
	}{
		{"n=1", 1, nil, "a", []int{1}},
		{"testcase1", 7, [][]int{[]int{0, 1}, []int{0, 2}, []int{1, 4}, []int{1, 5}, []int{2, 3}, []int{2, 6}},
			"abaedcd", []int{2, 1, 1, 1, 1, 1, 1}},
		{"testcase2", 4, [][]int{[]int{0, 1}, []int{1, 2}, []int{0, 3}}, "bbbb", []int{4, 2, 1, 1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numberOfNodes(tt.n, tt.edges, tt.labels)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("with input +wanted -got %s", diff)
			}
			t.Log("pass")
		})
	}
}
