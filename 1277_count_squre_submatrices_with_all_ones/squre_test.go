package square

import "testing"

func TestCountSqures(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]int
		exp    int
	}{
		{"empty matrix", nil, 0},
		{"all zeros", [][]int{[]int{0, 0}, []int{0, 0}, []int{0, 0}}, 0},
		{"all ones", [][]int{[]int{1, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1}}, 14},
		{"testcase1", [][]int{[]int{0, 1, 1, 1}, []int{1, 1, 1, 1}, []int{0, 1, 1, 1}}, 15},
		{"testcase2", [][]int{[]int{1, 0, 1}, []int{1, 1, 0}, []int{1, 1, 0}}, 7},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := countSquare(tt.matrix)
			if out != tt.exp {
				t.Fatalf("with input matrix:%v wanted %d but got %d", tt.matrix, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
