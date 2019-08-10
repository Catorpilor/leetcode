package oab

import "testing"

func TestMinTransfers(t *testing.T) {
	st := []struct {
		name  string
		input [][]int
		exp   int
	}{
		{"empty input", [][]int{}, 0},
		{"single input", [][]int{[]int{0, 1, 5}}, 1},
		{"lc testcase 1", [][]int{[]int{0, 1, 10}, []int{2, 0, 5}}, 2},
		{"lc testcase 2", [][]int{[]int{0, 1, 10}, []int{1, 0, 1}, []int{1, 2, 5}, []int{2, 0, 5}}, 1},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := MinTransfers(tt.input)
			if out != tt.exp {
				t.Fatalf("with input %v, wanted %d but got %d", tt.input, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
