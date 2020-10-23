package dna

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindRepeatedSeq(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  []string
	}{
		{"testcase1", "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		{"testcase2", "AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findRepeatedSeq(tt.s)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
