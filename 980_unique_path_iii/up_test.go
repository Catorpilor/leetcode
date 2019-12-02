package up

import "testing"

func TestUniquePath(t *testing.T) {
	st := []struct {
		name string
		grid [][]int
		exp  int
	}{
		{"empty grid", [][]int{}, 0},
		{"testcase1", [][]int{[]int{1, 0, 0, 0}, []int{0, 0, 0, 0}, []int{0, 0, 2, -1}}, 2},
		{"testcase2", [][]int{[]int{1, 0, 0, 0}, []int{0, 0, 0, 0}, []int{0, 0, 0, 2}}, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := uniquePath(tt.grid)
			if out != tt.exp {
				t.Fatalf("with grid: %v wanted %d but got %d", tt.grid, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
