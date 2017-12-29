package circle

import "testing"

func TestCircles(t *testing.T) {
	st := []struct {
		name string
		grid [][]int
		exp  int
	}{
		{"empty", [][]int{}, 0},
		{"single", [][]int{[]int{1}}, 1},
		{"test1", [][]int{[]int{1, 1, 0}, []int{1, 1, 0}, []int{0, 0, 1}}, 2},
		{"test2", [][]int{[]int{1, 1, 0}, []int{1, 1, 1}, []int{0, 1, 1}}, 1},
		{"failed1", [][]int{[]int{1, 0, 0, 1}, []int{0, 1, 1, 0}, []int{0, 1, 1, 1}, []int{1, 0, 1, 1}}, 1},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Circles(c.grid)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.grid)
			}
		})
	}
}
