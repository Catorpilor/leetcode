package hindex

import (
	"testing"
)

func TestGethIndex(t *testing.T) {
	st := []struct {
		name      string
		citations []int
		exp       int
	}{
		{"all identical", []int{1, 1, 1, 1}, 1},
		{"testcase1", []int{0, 1, 3, 4, 5}, 3},
		{"testcase2", []int{0, 1, 1, 1, 2, 3, 4}, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := getHIndex(tt.citations)
			if out != tt.exp {
				t.Fatalf("with input citations:%v wanted %d but got %d", tt.citations, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
