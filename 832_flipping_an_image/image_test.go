package flipping

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFlipAndInvertImage(t *testing.T) {
	st := []struct {
		name int
		A    [][]int
		exp  [][]int
	}{
		{"empty image", nil, nil},
		{"testcase1", [][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}}, [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}},
		{"testcase2", [][]int{[]int{1, 1, 0, 0}, []int{1, 0, 0, 1}, []int{0, 1, 1, 1}, []int{1, 0, 1, 0}}, [][]int{[]int{1, 1, 0, 0}, []int{0, 1, 1, 0}, []int{0, 0, 0, 1}, []int{1, 0, 1, 0}}},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := flipAndInvert(tt.A)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
