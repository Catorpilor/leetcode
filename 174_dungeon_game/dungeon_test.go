package dungeon

import "testing"

func TestMinHP(t *testing.T) {
	st := []struct {
		name string
		grid [][]int
		exp  int
	}{
		{"m=1", [][]int{[]int{-5, -3, -5}}, 14},
		{"n=1", [][]int{[]int{-2}, []int{-3}, []int{25}}, 6},
		{"testcaes1", [][]int{[]int{-2, -3, 3}, []int{-5, -10, 1}, []int{10, 30, -5}}, 7},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := calculateMinHP(tt.grid)
			if out != tt.exp {
				t.Fatalf("with input grid:%v wanted %d but got %d", tt.grid, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
