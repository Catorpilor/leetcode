package digits

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSeqDigits(t *testing.T) {
	st := []struct {
		name      string
		low, high int
		exp       []int
	}{
		{"testcase1", 10, 200, []int{11, 12, 23, 34, 45, 56, 67, 78, 89, 123}},
		{"testcase2", 100, 300, []int{123, 234}},
		{"testcase3", 1000, 13000, []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := seqDigits(tt.low, tt.high)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got), diff")
			}
			t.Log("pass")
		})
	}
}
