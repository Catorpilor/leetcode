package queue

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReconstructQueue(t *testing.T) {
	st := []struct {
		name   string
		people [][]int
		exp    [][]int
	}{
		{"testcase1", [][]int{[]int{7, 0}, []int{4, 4}, []int{7, 1}, []int{5, 0}, []int{6, 1}, []int{5, 2}},
			[][]int{[]int{5, 0}, []int{7, 0}, []int{5, 2}, []int{6, 1}, []int{4, 4}, []int{7, 1}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := reconstruct(tt.people)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("with input people:%v mismatch (-want +got):\n%s", tt.people, diff)
			}
			t.Log("pass")
		})
	}
}
