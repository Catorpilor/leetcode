package rot

import "testing"

func TestOrangeRotting(t *testing.T) {
	st := []struct {
		name string
		grid [][]int
		exp  int
	}{
		{"empty grid", nil, 0},
		{"no rotting ones", [][]int{[]int{0, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}, -1},
		{"testcase1", [][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}}, 4},
		{"testcase2", [][]int{[]int{2, 1, 1}, []int{0, 1, 1}, []int{1, 0, 1}}, -1},
		{"testcase3", [][]int{[]int{0, 2}}, 0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := orangeRotting(tt.grid)
			if out != tt.exp {
				t.Fatalf("with input grid wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
