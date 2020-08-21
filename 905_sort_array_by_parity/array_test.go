package array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortArrayByParity(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		exp  []int
	}{
		{"all even", []int{2, 4, 6, 8}, []int{2, 4, 6, 8}},
		{"all odds", []int{1, 3, 5}, []int{1, 3, 5}},
		{"testcase1", []int{3, 1, 2, 4}, []int{2, 4, 3, 1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := sortArrayByParity(tt.arr)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
