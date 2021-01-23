package matrix

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDiagonalSort(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]int
		exp    [][]int
	}{
		{"empty matrix", nil, nil},
		{"single row", [][]int{[]int{3, 2, 1}}, [][]int{[]int{3, 2, 1}}},
		{"testcase1", [][]int{[]int{3, 3, 1, 1}, []int{2, 2, 1, 2}, []int{1, 1, 1, 2}}, [][]int{[]int{1, 1, 1, 1}, []int{1, 2, 2, 2}, []int{1, 2, 3, 3}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := diagonalSort(tt.matrix)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pas")
		})
	}
}
