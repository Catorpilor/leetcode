package pancake

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPancakeSorting(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		exp  []int
	}{
		{"testcase1", []int{1, 2, 3}, []int{}},
		{"testcase2", []int{2, 1, 3, 4}, []int{1}},
		{"testcase3", []int{3, 2, 4, 1}, []int{3, 1, 3, 2}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := pancakeSorting(tt.arr)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted,+got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
