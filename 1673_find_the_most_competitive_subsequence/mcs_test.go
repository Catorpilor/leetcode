package mcs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMostCompetitive(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  []int
	}{
		{"nil slice", nil, 1, nil},
		{"n = k", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"k=1", []int{4, 3, 1, 5, 2}, 1, []int{1}},
		{"all identical", []int{1, 1, 1, 1, 1, 1}, 3, []int{1, 1, 1}},
		{"testcase1", []int{3, 5, 2, 6}, 2, []int{2, 6}},
		{"testcase2", []int{2, 4, 3, 3, 5, 4, 9, 6}, 4, []int{2, 3, 3, 4}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := mostCompetitive(tt.nums, tt.k)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted,+got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
