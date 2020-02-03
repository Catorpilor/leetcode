package mps

import "testing"

func TestMinPathSum(t *testing.T) {
	st := []struct {
		name string
		grid [][]int
		exp  int
	}{
		{"empty grid", [][]int{}, 0},
		{"all identical", [][]int{[]int{1, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1}}, 5},
		{"testcase1", [][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}, 7},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minPathSum(tt.grid)
			if out != tt.exp {
				t.Fatalf("with grid:%v wanted %d but got %d", tt.grid, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
