package interval

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRightInterval(t *testing.T) {
	st := []struct {
		name      string
		intervals [][]int
		exp       []int
	}{
		{"single one", [][]int{[]int{1, 2}}, []int{-1}},
		{"testcase1", [][]int{[]int{3, 4}, []int{1, 2}, []int{2, 3}}, []int{-1, 1, 0}},
		{"testcase2", [][]int{[]int{1, 4}, []int{2, 3}, []int{3, 4}}, []int{-1, 2, -1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := rightInterval(tt.intervals)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted,+got), %s", diff)
			}
			t.Log("pass")
		})
	}
}
