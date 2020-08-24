package verticies

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinimumVerticies(t *testing.T) {
	st := []struct {
		name string
		n    int
		edge [][]int
		exp  []int
	}{
		{"testcase1", 6, [][]int{[]int{0, 1}, []int{0, 2}, []int{2, 5}, []int{3, 4}, []int{4, 2}}, []int{0, 3}},
		{"testcase2", 5, [][]int{[]int{0, 1}, []int{2, 1}, []int{3, 1}, []int{1, 4}, []int{2, 4}}, []int{0, 2, 3}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minimumVerticies(tt.n, tt.edge)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
