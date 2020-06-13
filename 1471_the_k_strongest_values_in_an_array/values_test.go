package values

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetStrongestK(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		k    int
		exp  []int
	}{
		{"len(arr) = 1", []int{1}, 3, []int{1}},
		{"len(arr) = k", []int{1, 3, 4}, 3, []int{1, 3, 4}},
		{"testcase1", []int{1, 2, 3, 4, 5}, 2, []int{5, 1}},
		{"testcase2", []int{1, 1, 3, 5, 5}, 2, []int{5, 5}},
		{"testcase3", []int{6, -3, 7, 2, 11}, 3, []int{-3, 11, 2}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := getStrongest(tt.arr, tt.k)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
			t.Log("pass")
		})
	}
}
