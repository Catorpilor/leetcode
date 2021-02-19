package wk

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKWeakestRows(t *testing.T) {
	st := []struct {
		name string
		mat  [][]int
		k    int
		exp  []int
	}{
		{"testcase1", [][]int{[]int{1, 1, 0, 0, 0}, []int{1, 1, 1, 1, 0}, []int{1, 0, 0, 0, 0}, []int{1, 1, 1, 1, 1}}, 3, []int{2, 0, 3}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := kWeakestRows(tt.mat, tt.k)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
