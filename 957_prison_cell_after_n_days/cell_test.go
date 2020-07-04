package cell

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAfterNDays(t *testing.T) {
	st := []struct {
		name  string
		cells []int
		n     int
		exp   []int
	}{
		{"testcase1", []int{0, 1, 0, 1, 1, 0, 0, 1}, 7, []int{0, 0, 1, 1, 0, 0, 0, 0}},
		// {"testcase2", []int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000, []int{0, 0, 1, 1, 1, 1, 1, 0}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := afterNDays(tt.cells, tt.n)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("mismatch (-wanted,+got), :%s", diff)
			}
			t.Log("pass")
		})
	}
}
