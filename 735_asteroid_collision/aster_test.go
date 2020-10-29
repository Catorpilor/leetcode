package aster

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAsterCollision(t *testing.T) {
	st := []struct {
		name      string
		asteroids []int
		exp       []int
	}{
		{"single node", []int{1}, []int{1}},
		{"same direction", []int{-1, -2}, []int{-1, -2}},
		{"testcase1", []int{5, 10, -5}, []int{5, 10}},
		{"testcase2", []int{8, -8}, []int{}},
		{"testcase3", []int{10, 2, -5}, []int{10}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := asterCollision(tt.asteroids)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("(-wanted, +got) %s", diff)
			}
			t.Log("pass")
		})
	}
}
