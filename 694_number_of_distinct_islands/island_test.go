package island

import "testing"

func TestNumOfDistinctIslands(t *testing.T) {
	st := []struct {
		name string
		grid [][]int
		exp  int
	}{
		{"empty", [][]int{}, 0},
		{"single0", [][]int{[]int{0}}, 0},
		{"test1", [][]int{[]int{1, 1, 0, 0}, []int{1, 1, 0, 0}, []int{0, 0, 1, 1}, []int{0, 0, 1, 1}}, 1},
		{"test2", [][]int{[]int{1, 1, 0, 1, 1}, []int{1, 0, 0, 0, 0}, []int{0, 0, 0, 0, 1}, []int{1, 1, 0, 1, 1}}, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := NumOfDistinctIslands(c.grid)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.grid)
			}
		})
	}
}
