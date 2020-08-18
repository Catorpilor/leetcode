package number

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumsSameDiff(t *testing.T) {
	st := []struct {
		name string
		n    int
		k    int
		exp  []int
	}{
		{"testcase1", 1, 0, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"testcase2", 2, 1, []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}},
		{"testcase3", 3, 7, []int{181, 292, 707, 818, 929}},
		{"failed1", 2, 0, []int{11, 22, 33, 44, 55, 66, 77, 88, 99}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numsSameDiff(tt.n, tt.k)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
