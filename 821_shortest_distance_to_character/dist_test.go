package dist

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShortestDistance(t *testing.T) {
	st := []struct {
		name string
		s    string
		c    byte
		exp  []int
	}{
		{"single character", "a", 'a', []int{0}},
		{"all identical", "aaa", 'a', []int{0, 0, 0}},
		{"testcase1", "loveleetcode", 'e', []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		{"testcase2", "aaab", 'b', []int{3, 2, 1, 0}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := shortestDistance(tt.s, tt.c)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s")
			}
			t.Log("pass")
		})
	}
}
