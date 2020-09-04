package pl

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPartitionLabels(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  []int
	}{
		{"empty string", "", nil},
		{"single one", "a", []int{1}},
		{"testcase1", "ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"testcase2", "abafdafdjklifjkldasjklafida", []int{27}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := labels(tt.s)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
